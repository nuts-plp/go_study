package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func KafkaConsumer() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("failed to start consumer! err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("fail to get list of partition! err:%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { //遍历每个分区
		//针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d!  err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		//异步每个分区的信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("partition:%d、offset:%d、key:%v、value:%v\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)

	}
}
func main() {
	KafkaConsumer()
}
