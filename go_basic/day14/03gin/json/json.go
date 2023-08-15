package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type json struct {
	//`binding:"required"`修饰的字段，接收值若为空值，则报错   该字段为必须值
	username string `form:"username" json:"username" xml:"username" url:"username" binding:"required"`
	password string `form:"password" json:"password" xml:"password" url:"password"`
}

func main() {
	route := gin.Default()
	//json绑定
	route.POST("loginjson", func(context *gin.Context) {
		//声明接收的变量
		var js json
		//Bind()默认解析并绑定form格式
		//根据content-type自动推断
		//err:=context.Bind(&js)

		//将request中body数据，自动按照json格式解析到结构体
		if err := context.ShouldBindJSON(&js); err != nil {
			//返回错误信息
			//gin.H封装成json数据的工具
			context.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		//判断用户名密码是否正确
		if js.username != "root" && js.password != "admin" {
			context.JSON(http.StatusBadRequest, gin.H{"statue": "200"})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"statue": "304"})
		}

	})
	route.Run(":8900")
}
