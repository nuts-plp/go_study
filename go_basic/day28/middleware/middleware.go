package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//  中间件1和2是为了测试 中间件的执行顺序  c.Next()  以及中间件间数据的共享
func MiddleWare1(c *gin.Context) {
	fmt.Println("我是第一个中间件--1")
	// 中间件中共用数据，示例如下
	c.Set("name", "潘丽萍")

	c.Next()
	fmt.Println("我是第一个中间件---2")
}

func MiddleWare2(c *gin.Context) {
	fmt.Println("我是第二个中间件------1")
	value, exists := c.Get("name")
	if exists != true {
		fmt.Println("之错误！")
	}
	value, ok := value.(string)
	if !ok {
		fmt.Println("断言错误！")
	}
	fmt.Println(value)
	c.Next()
	fmt.Println("我是第二个中间件------2")
}

// 中间件3和4是为了测试c.Abort()
func MiddleWare3(c *gin.Context) {
	fmt.Println("我是第三个中间件--1")
	c.Next()
	fmt.Println("我是第三个中间件---2")
}

func MiddleWare4(c *gin.Context) {
	fmt.Println("我是第四个中间件------1")
	c.Abort()
	fmt.Println("我是第四个中间件------2")
}

// 中间件5是为了测试在中间件中开了协程要使用 c.Copy
func MiddleWare5(c *gin.Context) {
	fmt.Println("我是第一个中间件--1")
	// 如果中间件需要使用协程，使用 c.Copy 传递context副本

	cCp := c.Copy()
	go func() {
		fmt.Println(cCp.Request.URL.Path)
	}()
	fmt.Println("我是第一个中间件---2")
}
