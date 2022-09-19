package main

import (
	"go_basic/day13/04logagent/config"
	"go_basic/day13/04logagent/etcd"
	"go_basic/day13/04logagent/kafka"
	"go_basic/day13/04logagent/taillog"
	"gopkg.in/ini.v1"
	"sync"

	"fmt"
)

var (
	cfg = new(config.AppConfig)
	wg  sync.WaitGroup
)

//读取日志发送到kafka
//func run() {
//
//	//读取日志
//	for {
//		select {
//		case line := <-taillog.ReadLog():
//			//发送到kafka
//			kafka.SendToKafka(cfg.KafkaConfig.Topic, line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//
//	}
//}
func main() {
	//加载配置文件
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Printf("load ini failed!   err: %v\n", err)
		return
	}

	//1、初始化kafka连接
	err = kafka.Init(cfg.KafkaConfig.Address, cfg.KafkaConfig.ChanMaxSize)
	if err != nil {
		fmt.Printf("init kafka failed!   err: %v\n", err)
		return
	}
	fmt.Println("init kafka successed!")
	//2、初始化taillog连接
	//err = taillog.init(cfg.TaillogConfig.FieName)
	//if err != nil {
	//	fmt.Printf("init taillog failed! err: %v\n", err)
	//	return
	//}
	//fmt.Println("init tail succeed!")
	//3、初始化etcd连接
	err = etcd.Init(cfg.EtcdConfig.Address, 5)
	if err != nil {
		fmt.Printf("init etcd failed! err:%v\n", err)
		return
	}
	fmt.Println("etcd init succeed!")

	//run()
	//为了实现每个logagent都拉去自己独有的配置，所以要以自己的ip区分
	//ip, err := utils.GetOutboundIP()
	//if err != nil {
	//	panic(err)
	//}
	//etcdConfKey := fmt.Sprintf(cfg.EtcdConfig.Key, ip)
	//2.1、从etcd获取日志收集项的配置信息
	configs, err := etcd.GetConfig(cfg.EtcdConfig.Key)
	if err != nil {
		fmt.Printf("get etcd Config failed! err:%v\n", err)
		return
	}
	fmt.Println("get config from etcd succeed!", configs)
	//2.2、派一个哨兵去监视日志收集项的变化（实现logagent的热加载配置）

	for index, value := range configs {
		fmt.Printf("key:%v、value:%v\n", index, value)
	}
	//3、收集日志发往kafka
	taillog.Init(configs)
	//因为newConfChan访问了taskMger的NewConfChan，这个channel是在taillog.init()中初始化的
	newConfChan := taillog.NewConfChan() //从taillog中获取对外暴露的通道
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConfig.Key, newConfChan) //哨兵发现最新的配置信息会通知上面的那个通道
	wg.Wait()
	//for _, conf := range configs {
	//	taillog.NewTailTask(conf.Path, conf.Topic)
	//}

}
