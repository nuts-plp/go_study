package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var (
	wg sync.WaitGroup
)

func web1() {
	defer wg.Done()
	//1、创建路由
	//default使用了logger和recovery中间件
	r := gin.Default()
	//2、绑定路由，执行函数
	//gin.Context 封装了request和response

	r.SetTrustedProxies([]string{"192.168.0.6"})
	r.GET("/user/:name/*action", func(context *gin.Context) {
		//API参数通过Pare()获取    :name是获取这个位置的值，*action是获取从这个位置开始到结束的所有内容
		//获取如  http://localhost:8800/user/潘丽萍	/vbhjk/jjs
		name := context.Param("name")
		action := context.Param("action")
		context.String(http.StatusOK, fmt.Sprintf("你好! %s I miss you so much. %v", name, action))
	})

	//r.PUT("/xxxPUT")
	//r.POST("/xxxPUT", setting)

	//3、监听端口 默认8080

	r.Run(":8800") //端口可以自定义
}

func web2() {
	defer wg.Done()
	//1、创建路由
	//default使用了logger和recovery中间件
	r := gin.Default()
	//2、绑定路由，执行函数
	//gin.Context 封装了request和response

	r.SetTrustedProxies([]string{"192.168.0.6"})
	r.GET("/welcome", func(context *gin.Context) {

		//URL参数通过DefaultQuery()和Query()获取，前者可以设置默认值
		//获取如 http://localhost:8800/welcome?name=周小林
		name1 := context.DefaultQuery("name", "潘丽萍")
		context.String(http.StatusOK, fmt.Sprintf("你好! %s I miss you so much. ", name1))
	})

	//r.PUT("/xxxPUT")
	//r.POST("/xxxPUT", setting)

	//3、监听端口 默认8080

	r.Run(":8800") //端口可以自定义
}
func main() {
	wg.Add(1)
	//web1()
	web2()
	wg.Wait()
}
