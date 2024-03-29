package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "123.249.105.59:6379",
		DB:       0,
		Password: "950629",
		PoolSize: 10,
	})
	result, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("connect to redis failed err:%v", err)
		panic("wrong")
	}
	log.Print(result)
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetPrefix("[likui]")
	log.SetOutput(os.Stdout)
	test1(client)
}
func test1(client *redis.Client) {
	set := client.Set(context.Background(), "test", "不再喜欢", 0)
	log.Printf("operation:%v     result:%v", set.String(), set.Val())
	get := client.Get(context.Background(), "test")
	log.Printf("operation:%v     result:%v", get.String(), get.Val())
}
