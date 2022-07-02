package main
import(
	"github.com/Shopify/sarama"
	"fmt"
)

var(
	client sarama.SyncProducer //声明一个全局连接kafka的生产者client
)
//专门往kafka里写日志的模块

//初始化client
func Init(addrs []string)(err error){
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll //发送完数据需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true //成功交付的消息将在success channel返回

	// //构造一个消息
	// msg := &sarama.ProducerMessage{}
	// msg.Topic = "web_log"
	// msg.Value = sarama.StringEncoder("this is a test log!")

	//连接kafka
	client, err = sarama.NewSyncProducer(addrs,config)
	if err != nil {
		fmt.Printf("producer closed ! err:%v\n",err)
		return
	}
	fmt.Println("connect to kafka successed!")
	return


}

func SendToKafka(topic,message string){
	// //构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(message)

	//发送到kafka
	pid,offset,err:=client.SendMessage(msg)
	fmt.Println("xxx...")
	if err != nil{
		fmt.Printf("send message failed! err:%v\n",err)
		return
	}
	fmt.Printf("pid:%v   offset:%v\n",pid,offset)

}
func main(){
	err :=Init([]string{"127.0.0.1:9092"})
	if err != nil{
		fmt.Printf("init failed! err:%v\n",err)
		return
	}
	SendToKafka("web_log","this is a log")
}