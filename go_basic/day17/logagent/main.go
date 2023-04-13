package main

import (
	"fmt"
	"sync"

	"go_basic/day17/logagent/etcd"

	"go_basic/day17/logagent/conf"
	"go_basic/day17/logagent/kafka"
	"go_basic/day17/logagent/taillog"
	"gopkg.in/ini.v1"
)

var (
	appConf = new(conf.AppConf)
)

func main() {
	//使用ini加载配置文件
	//file, _ := ini.Load("./conf/config.ini")
	//fmt.Println(file.Section("kafka").Key("address").String())
	//fmt.Println(file.Section("kafka").Key("port").String())
	//fmt.Println(file.Section("kafka").Key("topic").String())
	//fmt.Println(file.Section("taillog").Key("filename").String())
	//fmt.Println(file.Section("").Key("app").String())

	//把配置文件映射到结构体
	//file, _ := ini.Load("./conf/config.ini")
	//app := new(conf.AppConf)
	//file.MapTo(app)
	//fmt.Println(app.Topic)
	//fmt.Println(app.Port)
	//fmt.Println(app.Address)
	//fmt.Println(app.Filename)

	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("load config file failed")
		return
	}
	file.MapTo(appConf)

	// 1、初始化kafka连接
	err = kafka.Init([]string{appConf.KafkaConf.Address}, appConf.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Println("kafka init failed")
		return
	}
	fmt.Println("初始化卡夫卡成功")

	//初始化etcd
	err = etcd.Init(appConf.EtcdConf.Address, appConf.EtcdConf.Timeout)
	if err != nil {
		fmt.Println("etcd init failed")
		return
	}
	//从etcd中获取配置项收集信息
	logEntryConf, err := etcd.GetConf(appConf.EtcdConf.Key)
	if err != nil {
		fmt.Println("get conf from etcd failed")
		return
	}
	fmt.Println("get conf from etcd successed")
	//初始化taillog
	taillog.Init(logEntryConf)
	//派一个哨兵去监视日志收集项的变化
	newConfChan := taillog.NewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(appConf.EtcdConf.Key, newConfChan)
	wg.Wait()
	// 2、打开日志文件
	//err = taillog.init(appConf.Filename)
	//if err != nil {
	//	fmt.Println("init taillog failed")
	//	return
	//}
	//fmt.Println("初始话taillog成功")
	//run()
}
