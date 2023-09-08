package router

import (
	"go_basic/day29/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func v3api(e *gin.Engine) {
	r := e.Group("/v2")
	r.GET("/uploadmultifilessamename", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uploadMultiFilesSameName.html", nil)
	})
	r.POST("/uploadmultifilessamename", controller.Admin{}.UploadMultiFileSameName)
}
