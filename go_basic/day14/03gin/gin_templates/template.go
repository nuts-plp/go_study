package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	//根据路径加载模板文件
	//route.LoadHTMLGlob("templates/*") //注意这个路径
	//根据文件名加载模板文件
	route.LoadHTMLFiles("templates/index.tmpl")
	route.GET("/index", func(context *gin.Context) {
		//根据文件名渲染
		//最终json将title替换
		context.HTML(200, "index.tmpl", gin.H{"title": "my lover! pan"})
	})
	route.Run(":8900")
}
