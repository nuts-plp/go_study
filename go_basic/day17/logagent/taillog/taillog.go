package taillog

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"go_basic/day17/logagent/kafka"
)

type TailTask struct {
	Path       string
	Topic      string
	instance   *tail.Tail
	context    context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (taskObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	taskObj = &TailTask{
		Path:       path,
		Topic:      topic,
		context:    ctx,
		cancelFunc: cancel,
	}
	taskObj.init()

	return
}

//从日志文件读日志的模块
func (t *TailTask) init() {
	var err error
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		MustExist: false,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		Poll:      true,
	}
	t.instance, err = tail.TailFile(t.Path, config)
	if err != nil {
		fmt.Println("tail tailObj failed")
		return
	}

	//直接采集日志发送到kafka
	go t.run()

}
func (t *TailTask) run() {
	for {
		select {
		case <-t.context.Done():
			fmt.Printf("taskObj %s_%s is over\n", t.Path, t.Topic)
			return
		case line := <-t.instance.Lines: //从tailObj通道中读取日志信息
			kafka.SendToChan(t.Topic, line.Text)
		}
	}
}
