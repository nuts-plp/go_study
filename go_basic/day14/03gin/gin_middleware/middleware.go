package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//中间件
//全局中间件和局部中间件

//定义中间件
func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行！")
		//设置变量到context的key中，可以通过get取
		c.Set("request:", "中间件")
		//执行中间件
		c.Next() //中间件最重要的一步
		statue := c.Writer.Status()
		fmt.Println("中间件执行完毕:", statue)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	route := gin.Default()
	//注册中间件
	route.Use(middleware()) //全局中间件的使用
	{                       //{}书写规范
		route.GET("/middleware", func(context *gin.Context) {
			//取值
			req, _ := context.Get("request:")
			fmt.Println("request:", req)
			//页面接受
			context.JSON(200, gin.H{"request": req})
		})
		//局部中间件的使用
		route.GET("/middleware1", middleware(), func(context *gin.Context) { //局部中间件的使用
			//取值
			req, _ := context.Get("request:")
			fmt.Println("request:", req)
			//页面接受
			context.JSON(200, gin.H{"request": req})
		})
	}
	route.Run(":8900")

}
