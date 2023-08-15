package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	engine := gin.Default()
	engine.Static("./static", "./static")

	//加载模板文件
	engine.LoadHTMLGlob("./templates/*")
	// 展示单文件上传
	v1api(engine)

	// 展示多文件同名上传
	v2api(engine)

	//展示多文件不同名上传
	v3api(engine)

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	engine.Run(":8000")
}
