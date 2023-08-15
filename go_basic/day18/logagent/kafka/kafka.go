package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var (
	client   sarama.SyncProducer
	confChan chan *logData
)

type logData struct {
	topic string
	data  string
}

//Init 通过传入kafka服务器地址，实现与服务器的连接
func Init(addr []string) (err error) {
	//设置关于kafka相关的设置
	config := sarama.NewConfig()
	//设置leader和所有follower确认才算完成
	config.Producer.RequiredAcks = sarama.WaitForAll
	//设置随机partitioner
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//设置消息发送的topic
	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("[kafka] create msg producer failed   error:", err)
		return
	}
	confChan = make(chan *logData)
	return
}

//SendToChan 包装消息发送至kafka通道中
func SendToChan(topic, data string) {
	logConfData := &logData{
		topic: topic,
		data:  data,
	}
	confChan <- logConfData

}

//sendToKafka 发送消息至kafka
func sendToKafka() {
	select {
	case logData := <-confChan:
		msg := &sarama.ProducerMessage{}
		msg.Topic = logData.topic
		msg.Value = sarama.StringEncoder(logData.data)
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("[kafka] send message to kafkaChan failed    error:", err)
			return
		}
		fmt.Printf("pid:%v   offset:%v\n", pid, offset)
	default:
		time.Sleep(time.Millisecond * 20)
	}
}
