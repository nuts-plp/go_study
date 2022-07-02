package main
import (
	"fmt"
	"time"
)
func main(){
	//时间定时器
	// ticker := time.Tick(time.Second)
	// for i := range ticker{
	// 	fmt.Println(i)
	// }
	//时间格式化
	now := time.Now()
	fmt.Println(now.Format("2006 01 02 03:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006-01-02 03:04:05.999"))
	//格式化时间部分
	fmt.Println(now.Format("3:04:05.000"))
	//格式化日期部分
	fmt.Println(now.Format("2006-01-02"))


	//解析字符串时间信息
	timeObj,err := time.Parse("2006/01/02 03:04:05","2022/06/10 03:33:09")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)

	loc,err := time.LoadLocation("Asia/Shanghai")
	if err != nil{
		fmt.Println(err)
		return
	}
	timeObj,err = time.ParseInLocation("2006/01/02 03:04:05","2022/06/10 03:33:09",loc)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}