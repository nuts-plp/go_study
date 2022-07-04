package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Printf("connect to etcd failed! err: %v\n", err)
		return
	}
	fmt.Println("connect to etcd succeeded!")
	defer cli.Close()
	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "zhou", "love")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed! err: %v\n", err)
		return
	}
	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"d:/k-v/tmp/etcd.log","topic":"etcd_log"},
               {"path":"d:/k_v/tmp/kafka.log","topic":"kafka_log"},
               {"path":"d:/k_v/tmp/redis.log","topic":"redis_log"}]`
	_, err = cli.Put(ctx, "/collect/config", value)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed ! err:%v\n", err)
		return
	}
}
