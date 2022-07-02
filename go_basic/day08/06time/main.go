package main
import(
	"fmt"
	"time"
)
//时间包
func main(){
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Date())
	fmt.Println()
	//时间戳
	fmt.Println(now.Unix())//秒级时间戳
	fmt.Println(now.UnixMicro())//微秒级时间戳
	fmt.Println(now.UnixNano())//纳秒级时间戳
	fmt.Println(now.UnixMilli())//微秒级时间戳
	fmt.Println("----------------------------")
	//将时间戳转换为时间对象
	secondsEastOfUTC := int((8*time.Hour).Seconds())
	beijing := time.FixedZone("beijing",secondsEastOfUTC)
	t := time.Date(2022,02,22,22,22,22,22,beijing)
	var(
		sec = t.Unix()
		micro = t.UnixMicro()
		milli = t.UnixMilli()
	)
	//将秒级时间戳转换为时间对象（第二个参数为不足一秒时的纳秒数）
	timeObj := time.Unix(sec,sec)
	fmt.Println(timeObj)
	timeObj = time.UnixMicro(micro)
	fmt.Println(timeObj)
	timeObj = time.UnixMilli(milli)
	fmt.Println(timeObj)
	//时间操作
	later := now.Add(2*time.Hour)
	fmt.Println(later)
}