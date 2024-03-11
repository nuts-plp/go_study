package routers

import "github.com/gin-gonic/gin"

func NewEngine() *gin.Engine {
	engine := gin.Default()
	en := handlers(engine)
	return en

}
