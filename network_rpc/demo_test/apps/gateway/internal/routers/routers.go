package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nuts/demo_test/apps/gateway/internal/logic"
	"github.com/nuts/demo_test/pkg/middleware"
)

func handlers(engine *gin.Engine) *gin.Engine {
	engine.GET("/ping", logic.Pong)
	engine.GET("/login", logic.Login)
	authEngine := engine.Group("/auth/v1")
	authEngine.Use(middleware.Auth, middleware.Core)
	{
		authEngine.POST("/hello", logic.Hello)
	}
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"msg": "source not found!"})
	})
	return engine
}
