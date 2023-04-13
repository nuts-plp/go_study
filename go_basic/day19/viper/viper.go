package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//设置默认值
	viper.SetDefault("filrdir", "./")
	//读取配置文件
	viper.SetConfigName("config") //配置文件名称 无拓展名
	viper.SetConfigType("yaml")   //如果配置文件中没有拓展名则需配置此项
	//viper.SetConfigFile("config.yaml")
	viper.AddConfigPath("/etcd/appname")  //查找配置文件所在路径
	viper.AddConfigPath("$HOME/.appname") //多次调用已填加多个搜索路径
	viper.AddConfigPath(".")              //还可以在工作目录中查找配置

	err := viper.ReadInConfig() //查找并读取配置文件
	if err != nil {
		panic(fmt.Errorf("fatal error config file:%s\n", err))
	}
	//实时监控配置文件的变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		//配置文件变更后会调用的回调函数
		fmt.Println("config file has changed file:", in.Name)
	})
	r := gin.Default()
	r.GET("/version", func(context *gin.Context) {
		context.String(200, viper.GetString("version"))
	})
	r.Run()
}
