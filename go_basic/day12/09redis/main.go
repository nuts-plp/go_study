package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisDB *redis.Client

//连接redis
func initDB() {
	//注意还有一种tls连接模式
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", //地址及端口号
		Password: "",               //密码
		DB:       0,                //数据库
		PoolSize: 20,               //连接池大小

	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		fmt.Printf("redis connection failed! error: %v", err)
		return
	}
	fmt.Println("redis connection successfully!")
}

func main() {
	initDB()

	//defer redisDB.Close()
}
