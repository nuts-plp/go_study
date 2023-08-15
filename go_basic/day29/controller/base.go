package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Base struct {
}

func (b Base) UploadSuccessed(c *gin.Context) {
	c.String(http.StatusOK, "图片上传成功！")

}

func (b Base) UploadFailed(c *gin.Context) {
	c.String(http.StatusOK, "图片上传失败！")

}
