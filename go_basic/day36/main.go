package main

import (
	"net/http"

	"github.com/micro/go-micro/web"

	"github.com/gin-gonic/gin"
)

func main() {
	{
		//service := web.NewService(func(o *web.Options) {
		//	o.Address = ":8000"
		//})
		// 或者
		//service:=web.NewService(web.Address(":8000))
		////使用go原生的api
		//service.HandleFunc("/hhh", func(writer http.ResponseWriter, request *http.Request) {
		//	writer.Write([]byte("afasdfsdgfa"))
		//
		//})
		//service.Run()
	}

	// 集成gin的api
	engine := gin.Default()
	engine.Handle("GET", "/", func(c *gin.Context) {
		c.String(http.StatusOK, "adskjfashkjd")
	})

	service := web.NewService(
		web.Address(":8000"),
		web.Handler(engine),
		// 注册到服务中心
		//web.Registry(consulReg)
	)

	service.Run()

}
