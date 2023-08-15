package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

type LogEntry struct {
	Path  string `json:path`
	Topic string `json:topic`
}

var (
	client *clientv3.Client
)

//Init 初始化etcd，连接etcd服务端 传入etcd服务端地址和连接超时时间限制
func Init(addr string, timeout int) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: time.Duration(timeout) * time.Second,
	})
	if err != nil {
		fmt.Println("[etcd] connect to etcd failed    error:", err)
		return
	}
	return
}

//GetConf 根据该方法传入的key，从etcd系统中获取对应的value值，并将值反序列化至一个切片中
func GetConf(key string) (logEntries []*LogEntry) {
	var err error
	var resp *clientv3.GetResponse
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err = client.Get(ctx, key)
	//fmt.Println(key, "-----------------------------")
	if err != nil {
		fmt.Println("[etcd] get conf failed  error:", err)
		return
	}
	cancel()
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntries)
		if err != nil {
			fmt.Println("[etcd] unmarshal logEntries failed    err:", err)
			return
		}
	}
	return
}

func WatchConf(key string, newConfChan chan<- []*LogEntry) {
	resp := client.Watch(context.Background(), key)
	for res := range resp {
		for _, event := range res.Events {
			var newConf []*LogEntry
			if event.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(event.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("[etcd] unmarshal newConf failed    error:", err)
					return
				}
			}
			newConfChan <- newConf
		}
	}
}
