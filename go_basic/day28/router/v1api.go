package router

import (
	"go_basic/day28/controller/articalController"
	"go_basic/day28/middleware"

	"github.com/gin-gonic/gin"
)

func v1api(r *gin.Engine) {
	e := r.Group("/v1", middleware.MiddleWare1, middleware.MiddleWare2, articalController.ArticalController{}.Edit)
	e.GET("/data", articalController.ArticalController{}.Edit)
}
