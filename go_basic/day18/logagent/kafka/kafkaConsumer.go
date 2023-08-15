package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"go_basic/day18/logagent/es"
	"time"
)

var (
	consumer      sarama.Consumer
	partitionList []int32
	pc            sarama.PartitionConsumer
)

func CreateConsumer(addr, topic string) (err error) {
	consumer, err = sarama.NewConsumer([]string{addr}, nil)
	if err != nil {
		err = fmt.Errorf("create consumer failed\n")
		return
	}
	partitionList, err = consumer.Partitions(topic)
	for partition := range partitionList {
		pc, err = consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			err = fmt.Errorf("create branch consumer failed")
			return
		}
		go func(partitionConsumer sarama.PartitionConsumer) {
			select {
			case msg := <-pc.Messages():
				es.SendMsgToChan(msg)
			default:
				time.Sleep(time.Millisecond * 50)
			}
		}(pc)
	}
	return
}
