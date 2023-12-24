package util

import (
	"github.com/gin-gonic/gin"
)

func GetUserIdByCookie(ctx *gin.Context) string {
	userId, _ := ctx.Cookie("user_id")
	return userId
}
