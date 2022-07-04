package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"strings"
)

//初始化生产者
func initProducer(str string) (err error) {
	config := nsq.NewConfig()
	_, err = nsq.NewProducer(str, config)
	if err != nil {
		fmt.Printf("create producer failed err: %v\n", err)
		return err
	}
	return nil
}

func main() {
	nsqAdress := "127.0.0.1:4150"
	err := initProducer(nsqAdress)
	if err != nil {
		fmt.Printf("initProducer failed err: %v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin) //从标准输入
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string from stdin failed! err:%v\n", err)
			continue
		}
		data = strings.TrimSpace(data)
		if err != nil {
			fmt.Printf("publish msg to nsq failed ! err:%v\n", err)
			continue
		}

	}
}
