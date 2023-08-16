package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello dockeruser!!!!")
	})
	err := engine.Run(":8080")
	fmt.Println(err)
}
