package main

import (
	"net/http"

	"github.com/gin-contrib/sessions/redis"

	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//设置cookie
	engine.GET("/", func(c *gin.Context) {
		c.SetCookie("email", "884639894", 3600, "/", "localhost", false, true)
		c.String(http.StatusOK, "设置cookie")
	})
	// 获取cookie
	engine.GET("/cookie", func(c *gin.Context) {
		email, err := c.Cookie("email")
		if err != nil {
			c.String(http.StatusOK, "获取cookie失败！")

		}
		c.String(http.StatusOK, "cookie="+email)
	})
	// 删除cookie  把cookie值置为空   或者把过期时间置为-1
	engine.GET("/delete", func(c *gin.Context) {
		c.SetCookie("email", "", -1, "/", "localhost", true, true)
		cookie, _ := c.Cookie("email")
		c.String(http.StatusOK, "cookie="+cookie)
	})
	// 设置二级域名共享cookie    a.sncot.top   b.sncot.top    注意路径设置为 /
	engine.GET("/delay", func(c *gin.Context) {
		c.SetCookie("hobby", "设置过期时间", 4, "/", ".sncot.top", false, true)
		c.String(http.StatusOK, "二级域名共享cookie")
	})

	/*
		session
		gin框架本身并没有集成session  可以使用第三方插件使用
		github.com/gin-contrib/sessions

	*/

	// 配置session中间件

	// 创建基于cookie的存储引擎，secret适用于加密的密钥
	{
		store := cookie.NewStore([]byte("secret"))
		// 配置session的中间件
		engine.Use(sessions.Sessions("mySession", store))
		engine.GET("/session", func(c *gin.Context) {
			// 设置session
			session := sessions.Default(c)
			session.Set("email", "我是第一个session")
			session.Save() //设置session时必须调用

			get := session.Get("email")
			s := get.(string)
			c.String(http.StatusOK, "设置了一个session|"+s)
		})
	}

	// 创建基于redis的session的存储方式
	{
		//创建基于redis的存储引擎‘
		store, _ := redis.NewStore(10, "tcp", "47.92.232.226:6379", "950629", []byte("sceret"))
		engine.Use(sessions.Sessions("mySession", store))
		engine.GET("/redis", func(c *gin.Context) {
			session := sessions.Default(c)
			//设置session的过期时间
			session.Options(sessions.Options{MaxAge: 3600 * 6}) //6hour
			session.Set("redis", "我是第一个redis存储的session")
			session.Save()
			get := session.Get("redis")
			s := get.(string)
			c.String(http.StatusOK, "redis存储的第一个session值+"+s)

		})
	}

	engine.Run(":80")

}
