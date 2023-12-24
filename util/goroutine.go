package util

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/vars"
)

func PanicHandlerWithContext(f func(*gin.Context)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		defer func() {
			if msg := recover(); msg != nil {
				vars.Log.Errorf("panic err:%v", msg)
			}
		}()
		f(ctx)
	}
}
