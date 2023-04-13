package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	filename := "./my.log"
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		Follow:    true,
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
	}
	file, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Println("tail file failed,", err)
		return
	}
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-file.Lines
		if !ok {
			fmt.Println("tail file close reopen,filename:%s\v", file.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
