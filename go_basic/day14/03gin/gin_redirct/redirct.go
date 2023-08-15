package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.SetTrustedProxies([]string{"192.168.0.6"})
	route.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.youdao.com/")
	})
	route.Run(":8900")
}
