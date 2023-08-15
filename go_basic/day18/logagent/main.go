package main

import (
	"fmt"
	"go_basic/day18/logagent/es"
	"go_basic/day18/logagent/taillog"
	"sync"

	"go_basic/day18/logagent/etcd"
	"go_basic/day18/logagent/kafka"

	"go_basic/day18/logagent/config"
	"gopkg.in/ini.v1"
)

func main() {
	// 0、首先获取配置信息
	var configEntry = new(config.Config)
	file, err := ini.Load("./config/conf.ini")
	if err != nil {
		fmt.Println("[main] get conf file failed")
		return
	}
	err = file.MapTo(configEntry)
	if err != nil {
		fmt.Println("[main] ini reflect failed   error:", err)
		return
	}
	// 1、首先初始化kafka、es、tailog和etcd
	err = kafka.Init([]string{configEntry.KafkaConf.Addr})
	if err != nil {
		fmt.Println("[main] kafka init failed")
		return
	}
	err = etcd.Init(configEntry.EtcdConf.Addr, configEntry.EtcdConf.Timeout)
	if err != nil {
		fmt.Println("[main] etcd init failed")
		return
	}
	err = es.Init(configEntry.EsConf.Addr, configEntry.EsConf.ChanMax, configEntry.EsConf.RoutineMax)
	if err != nil {
		fmt.Println("es init failed")
		return
	}
	// 2、从etcd中获取配置项
	logEntry := etcd.GetConf(configEntry.EtcdConf.Key)
	taillog.Init(logEntry)
	// 3、拍一个哨兵监视etcd中的配置
	newConfChan := taillog.GetNewConfChan()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go etcd.WatchConf(configEntry.EtcdConf.Key, newConfChan)
	wg.Wait()
	// 3、taillogMgr对从etcd获取的配置项进行遍历，分别启动一个任务协程
	taillog.Init(logEntry)
}
