package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/898989", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"hello": "asjdakjfas",
			"sfjs":  12,
		})

	})
	engine.Run(":8000")
}
