package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Println("connected to etcd failed")
		return
	}
	defer client.Close()
	resp := client.Watch(context.Background(), "mylover")
	for v := range resp {
		for _, event := range v.Events {
			fmt.Printf("Type:%v  key:%v   value:%v\n", event.Type, string(event.Kv.Key), string(event.Kv.Value))
		}

	}
}
