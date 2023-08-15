package main

import (
	"go_basic/day30/utils"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

var ma = map[string]bool{
	".jpg":  true,
	".png":  true,
	".gif":  true,
	".jpeg": true,
}

func main() {
	engine := gin.Default()
	engine.Static("/static", "./static")
	engine.LoadHTMLGlob("./templates/*")
	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	engine.POST("/post", func(c *gin.Context) {
		value := c.PostForm("username")
		file, err := c.FormFile("profile")
		if err != nil {
			c.String(http.StatusOK, "文件上传失败！")
		}
		// 获取文件上传那一刻的秒
		formatSeconds := utils.GetDay()
		ext := path.Ext(file.Filename)
		if _, ok := ma[ext]; !ok {
			c.String(http.StatusOK, "文件格式不合法！")
		}
		// 构建一个文件存储路径
		filedir := path.Join("./static/upload/", formatSeconds)
		err = os.MkdirAll(filedir, 0666)
		if err != nil {
			c.String(http.StatusOK, "创建文件夹失败！")
		}
		// 创建文件保存路径
		filepath := path.Join(filedir, utils.GetUnix()+ext)
		c.SaveUploadedFile(file, filepath)
		c.JSON(http.StatusOK, gin.H{
			"name":    value,
			"profile": filepath,
		})

	})
	engine.Run(":8000")

}
