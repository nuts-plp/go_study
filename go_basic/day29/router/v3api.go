package router

import (
	"go_basic/day29/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func v2api(e *gin.Engine) {
	r := e.Group("/v3")
	r.GET("/uploadmultifilesdifferentname", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uploadMultiFilesDifferentName.html", nil)
	})
	r.POST("/uploadmultifilesdifferentname", controller.Admin{}.UploadMultiFileDifferentName)
}
