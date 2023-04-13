package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端并校验
		if cookie, err := c.Cookie("dear pan"); err == nil {
			if cookie == "dear pan" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
	}
	//获取客户端cookie并校验
}
func main() {
	route := gin.Default()
	route.GET("/login", func(context *gin.Context) {
		//设置cookie
		context.SetCookie("dear pan", "亲爱的潘", 60, "/",
			"localhost", false, false)
		//返回信息
		context.String(200, "login sucessed!")
	})
	route.GET("/cookie", middleware(), func(context *gin.Context) {
		context.JSON(200, gin.H{"data": "lover"})

	})
	route.Run(":8900")
}
