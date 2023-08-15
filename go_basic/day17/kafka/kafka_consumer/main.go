package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	//获取消费者对象
	consumer, _ := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	//根据topic取到所有分区
	partitionList, _ := consumer.Partitions("web_log")
	fmt.Println(partitionList)
	//遍历所有的分区
	for partition := range partitionList {
		//针对每个分区创建一个对应的消费者
		pc, _ := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)

		defer pc.AsyncClose()
		//一部从每个分区消费消息
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("partition:%d   offset:%d    key:%s    value:%s   ", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	select {}
}
