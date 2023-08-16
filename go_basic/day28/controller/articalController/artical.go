package articalController

import (
	"go_basic/day28/controller/baseController"

	"github.com/gin-gonic/gin"
)

type ArticalController struct {
	baseController.BaseController
}

func (a ArticalController) Edit(c *gin.Context) {
	a.Success(c)
}
