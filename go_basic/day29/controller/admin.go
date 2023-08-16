package controller

import (
	"net/http"
	path2 "path"

	"github.com/gin-gonic/gin"
)

type Admin struct {
	Base
}

// 上传单个文件
func (receiver Admin) UplodaSingleFile(c *gin.Context) {
	value := c.PostForm("username")
	email := c.PostForm("email")
	file, err := c.FormFile("profile")
	if err != nil {
		receiver.UploadFailed(c)
	}
	path := path2.Join("./static/uploadfile", file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		receiver.UploadFailed(c)
	}
	c.JSON(http.StatusOK, gin.H{
		"username": value,
		"email":    email,
		"filepath": path,
	})
}

// 获取多个名称相同的文件
func (receiver Admin) UploadMultiFileSameName(c *gin.Context) {
	value := c.PostForm("username")
	email := c.PostForm("email")

	//获取相同名称的多个文件
	form, err := c.MultipartForm()
	if err != nil {
		receiver.UploadFailed(c)
	}
	files := form.File["profile[]"]
	var arr = make([]string, 0)
	for _, v := range files {
		path := path2.Join("./static/uploadfile", v.Filename)
		err := c.SaveUploadedFile(v, path)
		if err != nil {
			receiver.UploadFailed(c)
		}
		arr = append(arr, path)
	}
	c.JSON(http.StatusOK, gin.H{
		"username":        value,
		"email":           email,
		"same_name_files": arr,
	})

}

// 获取多个名称不同的文件
func (receive Admin) UploadMultiFileDifferentName(c *gin.Context) {
	value := c.PostForm("username")
	email := c.PostForm("email")
	file1, err := c.FormFile("profile1")
	if err != nil {
		receive.UploadFailed(c)
	}
	file2, err := c.FormFile("profile2")
	if err != nil {
		receive.UploadFailed(c)
	}
	path0 := path2.Join("./static/uploadfile", file1.Filename)
	path1 := path2.Join("./static/uploadfile", file2.Filename)
	arr := []string{path0, path1}
	c.JSON(http.StatusOK, gin.H{
		"username":             value,
		"email":                email,
		"different_name_files": arr,
	})
}
