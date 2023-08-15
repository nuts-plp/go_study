package main

//路由分组
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//  route 地址  github.com/julienschmidt/httprouter
//前缀树
//路由分组  routes group是为了管理一些相同的url
func main() {
	route := gin.Default()
	//DOS 命令 get curl http://localhost:8900/v1/login?name=于欢
	r1 := route.Group("/v1")
	//{}可有可无  书写规范
	{
		r1.GET("/login", login)
		r1.GET("/submit", submit)
	}
	//DOS 命令 post  curl http://localhost:8900/v2/submit -X POST
	r2 := route.Group("/v2")
	{
		r2.POST("/upload", upload)
		r2.POST("/submit", submit)
	}
	route.Run(":8900")

}
func login(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "潘丽萍")
	ctx.String(http.StatusOK, fmt.Sprintf("hello %s!  本方法是v1独有的", name))
}
func submit(ctx *gin.Context) {
	ctx.String(http.StatusOK, fmt.Sprintf("本方法是两个路由组共用的"))

}
func upload(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "周小林")
	ctx.String(http.StatusOK, fmt.Sprintf("hello %s !  本方法是v2独有的", name))
}
