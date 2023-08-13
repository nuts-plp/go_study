package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func v1api(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	v1.GET("/data", func(c *gin.Context) {

		c.String(http.StatusOK, "v1的get方法路由")
	})
	v1.POST("/data", func(c *gin.Context) {
		c.String(http.StatusOK, "v1的post方法路由")
	})
}
