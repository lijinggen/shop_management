package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/server/file_server"
	"github.com/shop_management/server/product_server"
	"github.com/shop_management/server/user_server"
	"net/http"
)

type serverFunc func(ctx *gin.Context) (interface{}, error)

func proxyFunc(in serverFunc) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		resp, err := in(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, err)
			return
		} else {
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func initRouter(engine *gin.Engine) {
	engine.Static("/static/", "./static")
	initUserRouter(engine)
	initUserTeam(engine)
	initFileApiRouter(engine)
	initProductApiRouter(engine)
}

func initUserRouter(engine *gin.Engine) {
	userServer := user_server.NewUserServer()
	engine.POST("/v1/api/user/login", proxyFunc(userServer.Login))
	engine.POST("/v1/api/user/register", proxyFunc(userServer.Register))
	engine.GET("/v1/api/user/profile", proxyFunc(userServer.GetUserProfile))

}

func initUserTeam(engine *gin.Engine) {
	userServer := user_server.NewUserTeamServer()
	engine.GET("/v1/api/user_team/sub_user_list", proxyFunc(userServer.SubUserList))
	engine.POST("/v1/api/user_team/add_sub_user", proxyFunc(userServer.AddSubUser))
	engine.POST("/v1/api/user_team/del_sub_user", proxyFunc(userServer.DelSubUser))
}

func initFileApiRouter(router *gin.Engine) {
	server := file_server.NewFileServer()
	router.POST("/v1/api/upload_file", server.Upload)
}

func initProductApiRouter(router *gin.Engine) {
	server := product_server.NewProductServer()
	router.POST("/v1/api/product/add", proxyFunc(server.Add))
}
