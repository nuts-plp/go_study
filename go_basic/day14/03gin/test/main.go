package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//err := InitDB()
	//if err != nil {
	//	panic(err)
	//}
	route := gin.Default()
	route.LoadHTMLGlob("./test/templates/*")
	route.GET("/", bookListHandler)
	route.Run(":8900")
}

func bookListHandler(c *gin.Context) {
	bookList, err := SelectBooks()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.HTML(200, "book_list.html", gin.H{
		"code": 0,
		"data": bookList,
	})

}
