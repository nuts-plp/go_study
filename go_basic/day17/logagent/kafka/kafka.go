package kafka

//专门往kafka里写日志的文件
import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer //声明一个全局的连接kafka的生产者client
	logDataChan chan *logData
)

func Init(addr []string, chanMixSize int) (err error) {
	//创建一个客户端实例
	config := sarama.NewConfig()
	//发送完消息要等待所有的结点确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机指认一个partitioner
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//成功交付的消息在success_chanel通道返回
	config.Producer.Return.Successes = true

	//连接kafka
	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("connect to kafka failed,err:", err)
		return
	}
	logDataChan = make(chan *logData, chanMixSize)
	//开启后台的goroutine去发送消息
	go sendToKafka()
	return
}

func SendToChan(topic, data string) {
	logData := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- logData
}

//sendToKafka 此方法用于异步从通道中取数据并发送至Kafka
func sendToKafka() {
	select {
	case logData := <-logDataChan:
		//创建一个消息
		msg := &sarama.ProducerMessage{}
		//设置topic
		msg.Topic = logData.topic
		msg.Value = sarama.StringEncoder(logData.data)
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("msg send failed,err:", err)
			return
		}
		fmt.Printf("pid:%v  offset:%v\n", pid, offset)
	default:
		time.Sleep(time.Millisecond * 50)
	}
}
