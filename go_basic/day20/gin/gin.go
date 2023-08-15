package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/hi", func(context *gin.Context) {
		context.String(200, "hello world!")
	})
	router.Run()
}
