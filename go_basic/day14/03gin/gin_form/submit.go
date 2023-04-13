package main

//取前端发来的form表单内容
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/form", func(context *gin.Context) {
		//获取表单信息
		//1、表单数设置默认值
		ty := context.DefaultPostForm("type", "alert")
		//接收其他值
		username := context.PostForm("username")
		password := context.PostForm("password")
		//接收复选框
		hobby := context.PostFormArray("hobby")
		context.String(http.StatusOK, fmt.Sprintf("type:%v、user:%s、password:%s、hobby:%v", ty, username, password, hobby))
	})
	r.Run(":8900")
}
