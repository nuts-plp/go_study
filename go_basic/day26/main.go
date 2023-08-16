package main

import (
	"encoding/xml"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Artical struct {
	Name    string
	Id      uint8
	Content string
}

func UnixToTime(timestamp time.Time) string {

	return timestamp.Format("2006-01-02-03:04")
}

type User struct {
	Name string `json:"name" form:"name" xml:"name"`
	Age  int    `json:"age" form:"age" age:"age"`
}

func main() {
	engine := gin.Default()
	//预制模板函数    注意把这个函数放在加载模板之前
	engine.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	// 配置静态web服务
	engine.Static("/static", "./static")

	// 加载模板，如果模板有分类，加载格式如下   注意：加载模板要在路由之前
	engine.LoadHTMLGlob("./templates/**/*")

	engine.GET("/data", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": 8091809182})
		context.Redirect(http.StatusOK, "/redirect")
	})
	engine.GET("/redirect", func(c *gin.Context) {
		c.String(http.StatusOK, "成功重定位")
	})
	engine.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "post请求")
	})
	engine.PUT("/put", func(c *gin.Context) {
		c.String(http.StatusOK, "put请求")
	})
	engine.PUT("/redirect", func(c *gin.Context) {
		c.String(http.StatusOK, "put重定位")
		c.Redirect(200, "dkasjdklasdal")
	})
	engine.PATCH("/data", func(c *gin.Context) {
		c.String(http.StatusOK, "asdasdas")
	})
	engine.DELETE("/", func(c *gin.Context) {
		c.Abort()
	})
	engine.GET("/json", func(c *gin.Context) {
		g := Artical{
			"啥子",
			89,
			"dhajshdakjs",
		}
		c.JSON(http.StatusOK, g)
	})
	engine.GET("/jsonp", func(c *gin.Context) {
		h := Artical{
			"hsjkda",
			79,
			"dasdasdasdasda",
		}
		c.JSONP(http.StatusOK, h)
	})
	engine.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"asdas": 8099,
			"8797":  "asdfjasdf",
		})
	})
	engine.GET("html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/dashboard.html", gin.H{
			"title": "你好，这是我的第一个文件",
			"price": 89,
			"date":  time.Now(),
		})
	})
	engine.GET("/html1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/dashboard.html", gin.H{
			"title": "你好，这是我的第一个文件",
			"price": 89,
			"date":  time.Now(),
		})
	})
	engine.GET("/html2", func(c *gin.Context) {
		k := Artical{
			"dasdasd",
			90,
			"dasdas",
		}
		v := []Artical{
			{"周小林", 22, "好久不见"},
			{"潘丽萍", 19, "喜欢你"},
		}
		c.HTML(http.StatusOK, "admin/dashboard", gin.H{
			"score": 89,
			"a":     k,
			"v":     v,
			"date":  time.Now(),
		})
	})
	// get 传参
	engine.GET("/getdata", func(c *gin.Context) {
		// 获取get传值
		//value := c.Query("name")
		//age := c.Query("age")
		user := &User{}
		err := c.ShouldBind(user)
		if err != nil {
			return
		}
		// 取值，如果没有设置默认值
		//query := c.DefaultQuery("name", "潘丽萍")
		c.JSON(http.StatusOK, gin.H{
			"name": user,
		})
	})

	//post  传参
	engine.POST("/postdata", func(c *gin.Context) {
		// post传值
		//value := c.PostForm("name")
		//form := c.PostForm("age")

		user := &User{}
		err := c.ShouldBind(user)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			//"name": value,
			"age": user,
		})
	})

	//获取 post xml数据
	engine.POST("/getxml", func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			return
		}
		user := &User{}
		err = xml.Unmarshal(data, user)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, user)
	})

	//动态路由传值     list/123      list/456
	engine.GET("/list/:cid", func(c *gin.Context) {
		param := c.Param("cid")
		c.String(http.StatusOK, param)
	})

	engine.Run(":8000")
}
