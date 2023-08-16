package router

import (
	"go_basic/day28/middleware"

	"github.com/gin-gonic/gin"
)

func Init() {
	engine := gin.Default()
	// 全局使用中间件
	engine.Use(middleware.MiddleWare5)
	v1api(engine)
	v2api(engine)
	engine.Run(":8000")
}
