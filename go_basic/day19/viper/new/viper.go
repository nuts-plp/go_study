package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		time.Sleep(time.Second * 5)
		context.String(200, "welcome to gin sever")
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listening %s\n", err)
		}
	}()
	//等待信号中断优雅的关闭服务器，为关闭服务器设置一个5秒的超时
	quit := make(chan os.Signal, 1) //创建一个接收信号的通道
	//kill默认会发送syscall.SIGTERM信号
	//kill -2会发送syscall.SIGINT 信号 我们常用的ctrl+c就是触发的SIGINI信号
	//kill -9会发送syscall.SIGKILL 信号 但是不能被捕获，所以不需要添加他
	//signal.Notify会把收到的syscall.SIGTERM或syscall.SIGINT信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //此处不会阻塞
	<-quit                                               //阻塞在此 当接收到上述两种信号时才会往下执行
	log.Println("shutdown server...")
	//创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	//五秒内优雅关闭服务 将未处理完的请求处理完再关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown ", err)
	}
	log.Println("server exit ...")
}
