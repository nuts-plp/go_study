package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func v2api(engine *gin.Engine) {
	v2 := engine.Group("/v2")
	v2.GET("/data", func(c *gin.Context) {
		c.String(http.StatusOK, "v2的get方法路由")
	})
	v2.POST("/data", func(c *gin.Context) {
		c.String(http.StatusOK, "v2的post方法的路由")
	})
}
