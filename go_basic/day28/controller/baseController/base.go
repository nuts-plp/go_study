package baseController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (b BaseController) Success(c *gin.Context) {
	c.String(http.StatusOK, "成功！")
}
func (b BaseController) Failed(c *gin.Context) {
	c.String(http.StatusNotFound, "失败！")

}
