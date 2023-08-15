package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	route := gin.Default()
	route.SetTrustedProxies([]string{"192.168.0.6"})
	//注意观察同步与异步的区别
	route.GET("/async", func(context *gin.Context) {
		//异步不要直接用context  创建一个context的副本
		acontext := context.Copy()
		go func() {
			time.Sleep(time.Second * 2)
			log.Println("异步请求：" + acontext.Request.URL.Path)
		}()

	})

	route.GET("/sync", func(context *gin.Context) {
		//同步请求 使用过原context
		time.Sleep(3 * time.Second)
		log.Println("同步请求:" + context.Request.URL.Path)
	})
	route.Run(":8900")
}
