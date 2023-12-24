package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron"
	"github.com/shop_management/util"
	"github.com/shop_management/vars"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"runtime"
	"time"
)

var SMApp = NewApp()

type App struct {
	GinEngine *gin.Engine
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	var err error
	// redis服务
	vars.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "120.24.169.86:4399",
		Password: "ljg098098",
		PoolSize: runtime.NumCPU() * 10,
	})
	// 日志
	initLogger()
	// 时间
	location, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = location
	// gin
	a.GinEngine = gin.Default()
	// 定时任务
	initCron()
	initInterceptor(a.GinEngine)
	initRouter(a.GinEngine)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	a.GinEngine.Use(cors.New(config))
	err = a.GinEngine.Run(":8080")
	if err != nil {
		log.Fatalf("running gin server failed, err:%v", err)
	}

}

func initLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	cfg := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:       "level",
			TimeKey:        "time",
			MessageKey:     "msg",
			CallerKey:      "caller",
			EncodeTime:     customTimeEncoder, // 使用自定义时间编码器
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"log/zap_info.txt"},
		ErrorOutputPaths: []string{"log/zap_error.txt"},
	}
	logger, err = cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	vars.Log = logger.Sugar()
}

func initCron() {
	c := cron.New()
	c.AddFunc("0 * * * *", func() {
		db, err := util.GetDB()
		if err != nil {
			return
		}
		defer func() {
			sqlDB, err := db.DB()
			if err == nil && sqlDB != nil {
				_ = sqlDB.Close()
			}
		}()

	})
	c.Start()
}
