package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var (
	redisDB *redis.Client
)

func initRedis() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "116.205.234.83:6379",
		Password: "root",
		DB:       0,
		PoolSize: 100,
	})
	s, err := redisDB.Ping().Result()
	if err != nil {
		err = fmt.Errorf("connect to redis failed")
		return
	}
	fmt.Println("connect to redis successfully")
	fmt.Println(s)
	return
}

//
////连接redis哨兵模式
//func initClient() {
//	redisDB = redis.NewFailoverClient(&redis.FailoverOptions{
//		MasterName:    "master",
//		SentinelAddrs: []string{"x.x.x.x:3234", "x.x.x.x:23456", "x.x.x.x:35423"},
//	})
//	_, err := redisDB.Ping().Result()
//	if err != nil {
//		err = fmt.Errorf("connect to redis failed")
//		return
//	}
//	return
//}
//
////集群模式
//func initCluster() {
//	redisObject := redis.NewClusterClient(&redis.ClusterOptions{
//		Addrs: []string{":7000", ":9000", "7863"},
//	})
//	_, err := redisObject.Ping().Result()
//	if err != nil {
//		fmt.Println("connect to redis failed")
//		return
//	}
//}

func test() {

	stat := redisDB.Set("潘丽萍", "mylover", time.Second)

	v, err := stat.Result()

	fmt.Println(stat.Name(), "|", stat.String(), "|", stat.Val(), "|", stat.Err(), "|", stat.Args(), "|", v, err, "|", stat)
	strCmd := redisDB.Get("潘丽萍")
	v, err = strCmd.Result()
	fmt.Println(strCmd, "|", strCmd.Val(), "|", strCmd.Err(), "|", strCmd.Args(), "|", strCmd.Name(), "|", strCmd.String(), "|", v, err)
	fmt.Println("--------------------------------------------")
	//redisDB.HGet()
	//redisDB.HMGet()
}
func main() {
	if err := initRedis(); err != nil {
		fmt.Println("init redis failed")
		return
	}
	test()
	defer redisDB.Close()
}
