package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func middle(c *gin.Context) {
	t := time.Now()
	c.Next()
	//统计时间
	t2 := time.Since(t)
	fmt.Println("time:", t2)
}
func main() {
	route := gin.Default()
	route.Use(middle)
	v1 := route.Group("/v1")
	{
		v1.GET("/m1", sleep1)
		v1.GET("/m2", sleep2)
	}
	route.Run(":8900")
}
func sleep1(c *gin.Context) {
	time.Sleep(time.Second * 3)
}

func sleep2(c *gin.Context) {
	time.Sleep(time.Second * 6)
}
