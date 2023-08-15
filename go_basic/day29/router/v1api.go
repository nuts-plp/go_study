package router

import (
	"go_basic/day29/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func v1api(e *gin.Engine) {
	r := e.Group("/v1")
	r.GET("/uploadsinglefile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uploadSingleFile.html", nil)
	})
	r.POST("/uploadsinglefile", controller.Admin{}.UplodaSingleFile)
}
