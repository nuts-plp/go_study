package middleware

import "github.com/gin-gonic/gin"

func Core(ctx *gin.Context) {
	ctx.Next()
}
