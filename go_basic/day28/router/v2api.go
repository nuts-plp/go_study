package router

import (
	"go_basic/day28/controller/articalController"
	"go_basic/day28/middleware"

	"github.com/gin-gonic/gin"
)

func v2api(r *gin.Engine) {
	//	使用全局中间件
	e := r.Group("/v2", middleware.MiddleWare3, middleware.MiddleWare4)
	e.GET("/data", articalController.ArticalController{}.Edit)
	e.POST("/data", articalController.ArticalController{}.Edit)
}
