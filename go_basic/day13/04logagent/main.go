package logagent

import (
	"time"

	"code.sncot.com/studygo/day13/logagent/kafka"
	"code.sncot.com/studygo/day13/logagent/taillog"
	"code.sncot.com/studygo/day13/logagent/config"
	// "github.com/hpcloud/tail"
	"gopkg.in/ini"

	"fmt"
)

var(
	cfg =new(config.AppConfig)
)

//读取日志发送到kafka
func run(){

	//读取日志
	for {
		select{
		case line := <-taillog.ReadLog():
			//发送到kafka
			kafka.SendToKafka(cfg.KafkaConfig.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}

		
	}
}
func main(){
	//加载配置文件

	err :=ini.MapTo("./config/config.ini")
	if err!=nil{
		fmt.Printf("load ini failed!   err: %v\n",err)
		return
	}

	cfg,err :=ini.Load("./config/config.ini")
	err =kafka.Init(cfg.KafkaConfig.Address)
	if err!=nil{
		fmt.Printf("init kafka failed!   err: %v\n", err)
		return
	}
	fmt.Println("init kafka successed!")
	err = taillog.Init(cfg.Taillog.FieName)
	if err!=nil{
		fmt.Printf("init taillog failed! err: %v\n", err)
		return
	}
	fmt.Println("init tail successed!")

	run()
}