package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

type logEntry struct {
	path  string `json:path`
	topic string `json:topic`
}

func main() {
	client, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, _ := client.Get(ctx, "xxx")
	cancel()
	var jso []*logEntry
	for _, v := range resp.Kvs {
		fmt.Println(string(v.Key), string(v.Value))
		json.Unmarshal(v.Value, &jso)
	}
	fmt.Println(jso)
}
