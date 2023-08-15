package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(context *gin.Context) {
		//获取单个文件
		//获取表单
		//file, _ := context.FormFile("file")
		//log.Println(file.Filename)
		////传到项目根目录，名字就用其本身的
		//context.SaveUploadedFile(file, file.Filename)
		////打印信息
		//context.String(http.StatusOK, fmt.Sprintf("%s uploaded!", file.Filename))

		//获取多个文件
		//1、获取表单
		form, _ := context.MultipartForm()
		//2、获取上传的图片
		pictures := form.File["files"]
		for _, file := range pictures {
			context.SaveUploadedFile(file, file.Filename)
			context.String(http.StatusOK, fmt.Sprintf("%s uploaded!", file.Filename))
		}

	})
	r.Run(":8900")
}
