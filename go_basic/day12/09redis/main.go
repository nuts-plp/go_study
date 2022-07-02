package main
import(
	"fmt"
	"github.com/go-redis/redis"
)
var redisDB *redis.DB
//连接redis
func initDB(){
	redisDB:=redis.NewClint(&redis.Options{
		addr:"127.0.0.1:6379"//地址及端口号
		PassWord:""//密码
		DB:0//数据库
		PoolSize:20//连接池大小

	})
	_,err:=redisDB.Ping().Result()
	if err != nil{
		fmt.Printf("redis connection failed! error: %v",err)
		return
	}
	fmt.Println("redis connection successfully!")
}
func main(){
	initDB()
	
}