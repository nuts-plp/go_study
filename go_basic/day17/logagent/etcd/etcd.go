package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	client *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

func Init(addr string, timeout int) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: time.Duration(timeout) * time.Second,
	})
	if err != nil {
		fmt.Println("etcd init failed")
		return
	}
	fmt.Println("etcd init successed")
	return
}

//从etcd中根据key获取配置项信息
func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	var resp *clientv3.GetResponse
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err = client.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Println("get conf failed")
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%v\n", ev.Key, string(ev.Value))
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Println("json.unmarshal etcd failed")
			return
		}
	}
	return

}

//作为哨兵监视配置信息的变化
func WatchConf(key string, newConfChan chan<- []*LogEntry) {
	ch := client.Watch(context.Background(), key)
	for resp := range ch {
		for _, event := range resp.Events {
			fmt.Printf("Type:%v    key:%v    value:%v\n", event.Type, event.Kv.Key, event.Kv.Value)
			//通知taillog.TaskMgr
			var newConf []*LogEntry
			if event.Type != clientv3.EventTypeDelete {
				//如果是删除操作  手动传递一个空的配置项
				err := json.Unmarshal(event.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("unmarshal failed")
					continue
				}
			}
			fmt.Println("get newConf", newConf)
			newConfChan <- newConf
		}
	}
}
