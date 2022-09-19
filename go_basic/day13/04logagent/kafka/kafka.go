package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	client      sarama.SyncProducer //声明一个全局连接kafka的生产者client
	logDataChan chan *logData
)

type logData struct {
	topic string
	data  string
}

//专门往kafka里写日志的模块

//初始化client
func Init(addrs string, maxSize int) (err error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel返回

	// //构造一个消息
	// msg := &sarama_kafka_producer.ProducerMessage{}
	// msg.Topic = "web_log"
	// msg.Value = sarama_kafka_producer.StringEncoder("this is a test log!")

	//连接kafka
	client, err = sarama.NewSyncProducer([]string{addrs}, config)
	if err != nil {
		fmt.Printf("producer closed ! err:%v\n", err)
		return
	}
	//初始化logDataChan
	logDataChan = make(chan *logData, maxSize)
	//开启一个后台的goroutine从通道中取数据发往kafka
	go sendToKafka()
	return

}

//给外部暴露的一个函数，该函数只把日志数据发送到通道中
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

//真正往kafka发送日志的函数
func sendToKafka() {

	for {
		select {
		case ld := <-logDataChan:
			// //构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)

			//发送到kafka
			pid, offset, err := client.SendMessage(msg)
			fmt.Println("xxx...")
			if err != nil {
				fmt.Printf("send message failed! err:%v\n", err)
				return
			}
			fmt.Printf("pid:%v   offset:%v\n", pid, offset)

		}
	}

}
