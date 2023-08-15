package es

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

var (
	client  *elastic.Client
	msgChan chan *sarama.ConsumerMessage
)

//Init 初始化es 即连接上es
func Init(addr string, chanMax, routineMax int) (err error) {
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}
	client, err = elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		err = fmt.Errorf("es init failed\n")
		return
	}
	msgChan = make(chan *sarama.ConsumerMessage, chanMax)
	for i := 0; i < routineMax; i++ {
		go sendMsgToEs()
	}
	return

}

//SendMsgToChan 发送消息到通道，实现异步的处理
func SendMsgToChan(msg *sarama.ConsumerMessage) {
	msgChan <- msg
}

func sendMsgToEs() {
	for {
		select {
		case msg := <-msgChan:
			put, err := client.Index().Index("student").BodyString(string(msg.Value)).Do(context.Background())
			if err != nil {
				fmt.Println("receive msg failed")
				continue
			}
			fmt.Printf("Indexed user %s to index %s, type %s\n", put.Id, put.Index, put.Type)
		default:
			time.Sleep(time.Microsecond * 30)
		}
	}
}
