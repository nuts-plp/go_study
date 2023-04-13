package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/cookie", func(context *gin.Context) {
		//获取客户端是否携带cookie
		cookie, err := context.Cookie("key_cookie")
		if err != nil {
			cookie = "Not set"
			//给客户端设置cookie
			//maxAge cookie有效时间
			//path  cookie所在目录
			//domain  域名
			//secure  是否只能通过https访问
			//httpOnly  是否允许js获取自己的cookie
			context.SetCookie("key_cookie", "my lover! pan", 60*10,
				"/", "http://www.sncot.com", false, false)
		}
		fmt.Println("cookie:", cookie)
	})
	route.Run(":8900")
}
