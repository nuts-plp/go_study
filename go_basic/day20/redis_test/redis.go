package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var (
	rdb *redis.Client

	Number string
)

func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	s, err := rdb.Ping().Result()
	if err != nil {
		err = fmt.Errorf("connect to redis failed")
		return
	}
	fmt.Println("connect to redis successfully")
	fmt.Println(s)
	return
}

//获取随机验证码
func getCode() string {
	rand.Seed(time.Now().UnixMicro())
	var code string
	for i := 0; i < 6; i++ {
		c := rand.Intn(10)
		code += strconv.Itoa(c)
	}
	return code
}

//每个手机每天只能发送三次 验证码放到redis中 设置过期时间
func getRedisCode(phoneNumber string) {
	var (
		countKey string = "VerifyCode" + phoneNumber + ":val"
		codeKey  string = "VerifyCode" + phoneNumber + ":code"
	)
	Number = phoneNumber
	val := rdb.Get(countKey).Val()
	fmt.Println("-----------------------1---------------")
	fmt.Println(val)
	if val == "" {
		rdb.SetXX(countKey, 1, 60*60*24)
	} else {
		count, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("parse val from string to int failed")
			panic("")
		}
		if count < 3 {
			rdb.Incr(countKey)
		} else {
			fmt.Println("your phone has send three times")
			return
		}

	}
	fmt.Println("-------------------------------------2----------------------------------")
	//发送的验证码要放到redis 过期时间2min
	sCode := getCode()
	rdb.SetXX(codeKey, sCode, 60*2)

}

//验证码的校验
func verifyCode() {
	codeKey := "VerifyCode" + Number + ":code"
	val := rdb.Get(codeKey).Val()
	var scan string
	_, err := fmt.Scan(&val)
	if err != nil {
		panic("Scan from terminal failed")
	}
	if scan == val {
		fmt.Println("pass~~~~")
	} else {
		fmt.Println("failed~~~")
	}
	defer rdb.Close()
}
func main() {
	//连接redis
	initRedis()
	//获取验证码
	getRedisCode("15629137870")
	//验证验证码
	verifyCode()
}
