package main

import "github.com/gin-gonic/gin"

func main() {
	c := gin.Default()
	c.GET("/", func(context *gin.Context) {
		context.String(200, "ok")
	})

	c.Run()
}
