package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.ClusterClient

//集群模式
func initCluster() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"47.103.141.125:7001", "47.103.141.125:7002", "47.103.141.125:7003", "47.103.141.125:7005", "47.103.141.125:7006"},
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("connect to redis failed")
		return
	}
}
func main() {
	initCluster()
	defer client.Close()
	client.Set("潘丽萍", "mylover", 10)
	cmd := client.Get("潘丽萍")
	fmt.Println(cmd.Val())
}
