package main

import (
	"fmt"
	"github.com/go-redis/redis" //李鬼
	"log"
	"os"
	// "github.com/redis/go-redis/v9"   正版
)

func initRedis() (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     "123.249.105.59:6379",
		Password: "950629",
		DB:       0,
		PoolSize: 100,
	})
	//client = redis.NewClusterClient(&redis.ClusterOptions{
	//	Addrs:    []string{"123.249.105.59"},
	//	Password: "950629",
	//	PoolSize: 10,
	//})
	s, err := client.Ping().Result()
	if err != nil {
		log.Println("connect to redis failed!")
		log.Printf("ping.result   s:%s", s)
		panic("create redis client failed!")
	}
	log.Println("connect to redis successfully")
	return client, err
}

// //连接redis哨兵模式
//
//	func initClient() {
//		redisDB = redis.NewFailoverClient(&redis.FailoverOptions{
//			MasterName:    "master",
//			SentinelAddrs: []string{"x.x.x.x:3234", "x.x.x.x:23456", "x.x.x.x:35423"},
//		})
//		_, err := redisDB.Ping().Result()
//		if err != nil {
//			err = fmt.Errorf("connect to redis failed")
//			return
//		}
//		return
//	}
//
// //集群模式
//
//	func initCluster() {
//		redisObject := redis.NewClusterClient(&redis.ClusterOptions{
//			Addrs: []string{":7000", ":9000", "7863"},
//		})
//		_, err := redisObject.Ping().Result()
//		if err != nil {
//			fmt.Println("connect to redis failed")
//			return
//		}
//	}
type Student struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int8   `json:"gender"`
	Phone  string `json:"phone"`
}

func test(client *redis.Client) {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime)
	log.SetPrefix("[redis]")
	/************************************************
	 ******************测试返回值*Statuscmd和*StringCmd ***
	 	**************顺便测试setnx和set和get *********/
	//stat := client.Set("潘丽萍", "mylover", time.Second)
	//
	//v, err := stat.Result()
	//
	//fmt.Println("name:", stat.Name())
	//fmt.Println("string:", stat.String())
	//fmt.Println("val:", stat.Val())
	//fmt.Println("err:", stat.Err())
	//fmt.Println("args:", stat.Args())
	//fmt.Println("result.v:", v)
	//fmt.Println("result.err:", err)
	//fmt.Println("--------------------------------------------")
	//statv := client.Get("潘丽萍")
	//v, err = statv.Result()
	//fmt.Println("name:", statv.Name())
	//fmt.Println("string:", statv.String())
	//fmt.Println("val:", statv.Val())
	//fmt.Println("err:", statv.Err())
	//fmt.Println("args:", statv.Args())
	//fmt.Println("result.v:", v)
	//fmt.Println("result.err:", err)
	///********************************************************
	//*******测试incr、incrby、decr、decrby***************************
	// ***********************************************************/
	//incr := client.Incr("潘丽萍")
	//fmt.Println("string:", incr.String())
	//decr := client.Decr("潘丽萍")
	//fmt.Println("string:", decr.String())
	//r := client.Set("l", 8, 0)
	//log.Printf("[redis] opteration:%v result:%v\n", r.String(), r.Val())
	//re := client.Incr("l")
	//log.Printf("[redis] opteration:%v result:%v\n", re.String(), re.Val())
	//res := client.IncrBy("l", int64(89))
	//log.Printf("[redis] opteration:%v result:%v\n", res.String(), res.Val())
	//resu := client.Decr("l")
	//log.Printf("[redis] opteration:%v result:%v\n", resu.String(), resu.Val())
	//resul := client.DecrBy("l", int64(50))
	//log.Printf("[redis] opteration:%v result:%v\n", resul.String(), resul.Val())
	/*********************************************************
	**************存储序列化后的结构体、获取后反序列化****************
	********************************************************* */
	//s := &Student{
	//	Id:     2131874897134812734,
	//	Name:   "afsdfasdf",
	//	Gender: 0,
	//	Phone:  "134234324234",
	//}
	//bytes, err := json.Marshal(s)
	//if err != nil {
	//	log.Fatalf("marshal student failed err:%v\n", err)
	//}
	//nx := client.SetNX("student", bytes, time.Second)
	//fmt.Println("string:", nx.String())
	//get := client.Get("student")
	//fmt.Println("string:", get.String())
	//stu := Student{}
	//err = json.Unmarshal([]byte(get.Val()), &stu)
	//if err != nil {
	//	log.Printf("unmaeshal stu failed err:%v\n", err)
	//}
	//fmt.Println("stu:", stu)
	/*******************************************************
	**********************list的命令rpush、rpop、lpush、lpop、lrange、lindex**********
	 ******************************************************************************/
	//rpush := client.RPush("lover", "潘")
	//log.Printf("opteration:%v    result:%v", rpush.String(), rpush.Val())
	//lpush := client.LPush("lover", "丽")
	//_ = client.LPush("lover", "萍")
	//log.Printf("opteration:%v    result:%v", lpush.String(), lpush.Val())
	//rpop := client.RPop("lover")
	//log.Printf("opteration:%v    result:%v", rpop.String(), rpop.Val())
	//lpop := client.LPop("lover")
	//log.Printf("opteration:%v    result:%v", lpop.String(), lpop.Val())
	//for i := 0; i < 10; i++ {
	//	_ = client.LPush("number", i)
	//}
	//log.Printf("opteration:%v    result:%v", rpush.String(), rpush.Val())
	//lRange := client.LRange("number", 1, 6)
	//log.Printf("opteration:%v    result:%v", lRange.String(), lRange.Val())
	//lIndex := client.LIndex("number", 0)
	//log.Printf("opteration:%v    result:%v", lIndex.String(), lIndex.Val())
	/***************************************************************
	***********************集合set的操作 sadd、scard、smembers、sismember********
	************************************************************************
	 */
	//sAdd := client.SAdd("set", "1", 2, "asfhasdfsadf", 7809018)
	//log.Printf("opteration:%v    result:%v", sAdd.String(), sAdd.Val())
	//sCard := client.SCard("set")
	//log.Printf("opteration:%v    result:%v", sCard.String(), sCard.Val())
	//sMembers := client.SMembers("set")
	//log.Printf("opteration:%v    result:%v", sMembers.String(), sMembers.Val())
	//sIsMember := client.SIsMember("set", 1)
	//log.Printf("opteration:%v    result:%v", sIsMember.String(), sIsMember.Val())
	/***********************************************************************
	*******************hash散列 适合存储对象 操作 hset、hget、hgetall、hdel*******
	***********************************************************************
	 */
	//hSet := client.HSet("潘丽萍", "name", "潘丽萍")
	//log.Printf("opteration:%v    result:%v", hSet.String(), hSet.Val())
	//_ = client.HSet("潘丽萍", "gender", 0)
	//_ = client.HSet("潘丽萍", "character", "作")
	//hGet := client.HGet("潘丽萍", "character")
	//log.Printf("opteration:%v    result:%v", hGet.String(), hGet.Val())
	//hGetAll := client.HGetAll("潘丽萍")
	//log.Printf("opteration:%v    result:%v", hGetAll.String(), hGetAll.Val())
	//hDel := client.HDel("潘丽萍", "gender")
	//log.Printf("opteration:%v    result:%v", hDel.String(), hDel.Val())
	//hGetAll = client.HGetAll("潘丽萍")
	//log.Printf("opteration:%v    result:%v", hGetAll.String(), hGetAll.Val())
	/************************************************************************
	**************有序结合zset 、成员唯一、score可重复 操作zadd、zrange、zrem**************
	********************************************************************************
	 */
	//zAdd := client.ZAdd("ll", redis.Z{Member: "潘丽萍", Score: 60}, redis.Z{Member: "周小林", Score: 80})
	//log.Printf("opteration:%v    result:%v", zAdd.String(), zAdd.Err())
	//zRange := client.ZRange("ll", 50, 70)
	//log.Printf("opteration:%v    result:%v", zRange.String(), zRange.Err())
	//zRem := client.ZRem("ll", "周小林")
	//log.Printf("opteration:%v    result:%v", zRem.String(), zRem.Err())
	/*********************************************************************************
	**********************hyperloglogs 基数统计  操作 pfadd、pfcount、pfmerge*********************************
	*********************************************************************************
	 */
	//pfAdd := client.PFAdd("cv", "asdfasd", "fsdfs", "asfasdfwe")
	//client.PFAdd("key", "asdas", "ada", 8, "dasda")
	//log.Printf("opteration:%v    result:%v", pfAdd.String(), pfAdd.Val())
	//pfCount := client.PFCount("cv")
	//log.Printf("opteration:%v    result:%v", pfCount.String(), pfCount.Val())
	//pfMerge := client.PFMerge("k3", "cv", "key")
	//log.Printf("opteration:%v    result:%v", pfMerge.String(), pfMerge.Val())
	/***************************************************************************
	************Bitmap位图 打卡、未打卡、登录、未登录   操作 setbit、getbit******
	*****************************************************************************
	 */
	//setBit := client.SetBit("sign", 0, 1)
	//log.Printf("opteration:%v    result:%v", setBit.String(), setBit.Val())
	//getBit := client.GetBit("sign", 0)
	//log.Printf("opteration:%v    result:%v", getBit.String(), getBit.Val())
	//getBit = client.GetBit("sign", 1)
	//log.Printf("opteration:%v    result:%v", getBit.String(), getBit.Err())
	/********************************************************************************
	**********************geospatial 地理位置  geoadd（添加）、geopos（获取）、************
	************geodist（如果不存在，返回空）、georadius（附近的人）、georadiusbymember******
	*************（显示指定成员一定半径其他成员）、geohash（返回11个字符的hash字符串）**********
	*********************************************************************************
	 */
	//geoAdd := client.GeoAdd("china:beijing" )
	//log.Printf("opteration:%v    result:%v", geoAdd.String(), geoAdd.Val())
	//geoDist := client.GeoDist()
	//log.Printf("opteration:%v    result:%v", geoDist.String(), geoDist.Val())
	//geoPos := client.GeoPos()
	//log.Printf("opteration:%v    result:%v", geoPos.String(), geoPos.Val())
	//geoRadius := client.GeoRadius()
	//log.Printf("opteration:%v    result:%v", geoRadius.String(), geoRadius.Val())
	//geoRadiusByMember := client.GeoRadiusByMember()
	//log.Printf("opteration:%v    result:%v", geoRadiusByMember.String(), geoRadiusByMember.Val())
	//geoHash := client.GeoHash()
	//log.Printf("opteration:%v    result:%v", geoHash.String(), geoHash.Val())
	/**********************************************************************************
	**********stream  stream   操作 xadd、xtrim、xdel、xlen、xrange、xrevrange、xread*******
	*********************************************************************************
	 */
	////*******错误示例**************/////////
	//xAdd := client.XAdd(&redis.XAddArgs{
	//	Stream: "ll",
	//	Values: map[string]interface{}{"name": "潘丽萍", "age": "23", "gender": "0"},
	//})
	//log.Printf("opteration:%v    result:%v", xAdd.String(), xAdd.Val())
	//xTrim := client.XTrim("ll", 64)
	//log.Printf("opteration:%v    result:%v", xTrim.String(), xTrim.Val())
	//xDel := client.XDel("ll", "gender")
	//log.Printf("opteration:%v    result:%v", xDel.String(), xDel.Val())
	//xLen := client.XLen("ll")
	//log.Printf("opteration:%v    result:%v", xLen.String(), xLen.Val())
	//xRange := client.XRange("ll", "-", "+")
	//log.Printf("opteration:%v    result:%v", xRange.String(), xRange.Val())
	//revRange := client.XRevRange("ll", "-", "+")
	//log.Printf("opteration:%v    result:%v", revRange.String(), revRange.Val())
	//xRead := client.XRead(&redis.XReadArgs{
	//	Count:   2,
	//	Streams: []string{"ll"},
	//})
	//log.Printf("opteration:%v    result:%v", xRead.String(), xRead.Val())
	/*******************************************************************************
	********************** publish/subscribe  操作 publish、subscribe、psubscribe、punsubscribe

	 */
	//ch := make(chan os.Signal)
	//signal.Notify(ch, os.Interrupt, os.Kill)
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		publish1 := client.Publish("kl.gh.vb", "hi"+strconv.Itoa(i))
	//		publish2 := client.Publish("kl.gh.vc", "hello"+strconv.Itoa(i))
	//		time.Sleep(time.Second)
	//		r1, _ := publish1.Result()
	//		r2, _ := publish2.Result()
	//
	//		log.Printf("opteration:%v    result:%v    result:%v", publish1.String(), publish1.Val(), r1)
	//		log.Printf("opteration:%v    result:%v    result:%v", publish2.String(), publish1.Val(), r2)
	//
	//	}
	//}()

	//go func() {
	//	subscribe := client.Subscribe("kl.gh.vb")
	//	defer subscribe.Close()
	//	//receive, _ := subscribe.Receive()
	//	ch := subscribe.Channel()
	//	for {
	//		c := <-ch
	//		fmt.Println("|", c, "|", c.Payload)
	//		//fmt.Println(receive)
	//	}
	//
	//}()
	//go func() {
	//	subscribe := client.PSubscribe("kl.gh.v?")
	//	cha := subscribe.Channel()
	//	for i := 0; i < 10000; i++ {
	//		msg := <-cha
	//		fmt.Println(msg, "|", msg.Payload)
	//		if i > 5 {
	//			_ = subscribe.Unsubscribe("kl.gh.vb")
	//		}
	//	}
	//
	//}()
	//<-ch
	//log.Printf("opteration:%v   value:%v ", subscribe.String(),)
	/**************************************************************************
	********************事务 操作 multi、exec、discard、watch、unwatch********************
	*************************************************************************
	 */
	//ch := make(chan struct{})
	//go func() {
	//	ch <- struct{}{}
	//	_ = client.Watch(func(tx *redis.Tx) error {
	//		log.Printf("bulabula  changed")
	//		return nil
	//	})
	//}()
	//<-ch
	//// 开启事务
	//pipe := client.TxPipeline()
	//set := pipe.Set("bulabula", "sdafsqfaf", time.Second*5)
	//log.Printf(" opterator:%v    result:%v ", set.String(), set.Val())
	//
	//pipe.Set("sa", "sdadsad", time.Second)
	//// 结束事务
	//_ = pipe.Discard()
	////提交事务
	//exec, _ := pipe.Exec()
	//log.Printf("%v", exec)

}
func main() {
	client, err := initRedis()
	if err != nil {
		fmt.Println("init redis failed")
		return
	}
	test(client)

	defer client.Close()
}
