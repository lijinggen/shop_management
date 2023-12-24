package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/redis/user_redis"
	"github.com/shop_management/sm_error"
	"github.com/shop_management/sm_error/error_code"
	"github.com/shop_management/util"
	"github.com/shop_management/vars"
	"gorm.io/gorm"
	"net/http"
	"runtime/debug"
)

func initInterceptor(engine *gin.Engine) {
	// panic
	engine.Use(func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				vars.Log.Errorf("panic,err:%v", err)
				context.JSON(http.StatusInternalServerError, sm_error.NewHttpError(error_code.PanicErr))
			}
		}()
		context.Next()
	})
	// db拦截器
	engine.Use(func(context *gin.Context) {
		db, err := util.GetDB()
		if err != nil {
			context.Abort()
		}
		if context.Request.URL.Path == "" {
		} else {
			defer func() {
				if db, ok := context.Get(vars.DbMetadataName); ok {
					if db, ok := db.(*gorm.DB); ok {
						sqlDB, err := db.DB()
						if err == nil && sqlDB != nil {
							_ = sqlDB.Close()
						}
					}
				}
			}()
		}
		context.Set(vars.DbMetadataName, db)
		context.Next()
	})
	// 跨域
	engine.Use(func(c *gin.Context) {
		allowHost := "http://localhost:5173"
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowHost) // 允许特定的域
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")          // 预检请求的有效期，单位为秒
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // 允许携带认证信息
		if c.Request.Method == "OPTIONS" {
			// 预检请求直接返回200
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
	// token 或 secret校验
	engine.Use(func(context *gin.Context) {
		if context.Request.URL.Path == "/v1/api/user/login" ||
			context.Request.URL.Path == "/v1/api/user/modify_password" ||
			context.Request.URL.Path == "/v1/api/user/register" ||
			context.Request.URL.Path == "/v1/api/sms_record/receive_report" {
			context.Next()
			return
		}
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("user_id")
		err := user_redis.CheckToken(context, token, userId)
		if err != nil {
			context.JSON(http.StatusOK, err)
			context.Abort()
			return
		}
		context.Next()
	})
}
