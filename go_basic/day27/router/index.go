package router

import "github.com/gin-gonic/gin"

func Init() {
	engine := gin.Default()
	v1api(engine)
	v2api(engine)
	engine.Run(":8000")
}
