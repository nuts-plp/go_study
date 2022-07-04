package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
)

// Init 初始化
func Init(addrs string, timeout int) (err error) {

	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addrs},
		DialTimeout: time.Duration(timeout) * time.Second,
	})

	if err != nil {
		fmt.Printf("connect to etcd failed! err: %v\n", err)
		return
	}
	fmt.Println("connect to etcd succeeded!")

	return
}

// LogEntry 需要收集日志的配置信息
type LogEntry struct {
	Path  string `json:"path"`  //日志存放的路径
	Topic string `json:"topic"` //日志要发往kafka的topic
}

//GetConfig 从etcd中根据key获取配置项
func GetConfig(key string) (configs []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed! err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		//fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &configs)
		if err != nil {
			fmt.Printf("unmarshal config failed! err:%v\n", err)
			return
		}
	}
	return
}

//WatchConf 监视配置的变化
func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)
	//尝试从通道取值（监视的信息）
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v  key:%v   value:%v \n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			//通知taillog.taskMger
			//1、先判断操作的类型
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				//如果是删除操作,手动传递一个空的配置项
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("unmarshal failed ! err:%v\n", err)
					continue
				}
			}

			fmt.Printf("get new conf! %v\n", newConf)
			newConfCh <- newConf
		}
	}
}
