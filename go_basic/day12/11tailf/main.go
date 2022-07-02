package main

import (
	"fmt"
	"time"  
	"github.com/hpcloud/tail"
)

func main(){
	fileName :="./my.log"
	config :=tail.Config{
		ReOpen: true,		//重新打开
		Follow: true,		//是否跟随
		Location:&tail.SeekInfo{Offset:0,Whence:2},//从文件的哪个地方开始读
		MustExist:false,		//文件不存在不报错
		Poll:true,
	}
	tail,err := tail.TailFile(fileName,config)
	if err != nil {
		fmt.Printf("tail file failed! err: %v\n", err)
		return
	}
	// var (
	// 	line *tail.Lines
	// 	ok bool
	// )
	for{
		line,ok :=<-tail.Lines
		if ok{
			fmt.Printf("tail file close reopen,filename: %s\n", fileName)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:",line.Text)

	}
}