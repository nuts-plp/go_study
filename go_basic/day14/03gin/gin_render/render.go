package main

//几种渲染方式
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()
	//1、json格式
	r.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"姓名": "周小林"})
	})

	//2、结构体格式
	r.GET("/struct", func(context *gin.Context) {
		var msg struct {
			name string
			age  int
		}
		msg.name = "牛敏"
		msg.age = 24
		context.JSON(http.StatusOK, msg)
	})
	//3、xml
	r.GET("/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"姓名": "潘丽萍"})
	})
	//4、YAML响应
	r.GET("/yaml", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"姓名": "于欢"})
	})
	//5、protobuf格式    谷歌开发的搞笑存储读取格式
	//如果自己创建一个存储格式，应该是怎样的？
	r.GET("/protobuf", func(context *gin.Context) {
		resp := []int64{1, 2, 3}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  resp,
		}
		context.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8900")
}
