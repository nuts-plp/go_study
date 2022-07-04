package taillog

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"go_basic/day13/04logagent/kafka"
)

var (
	tailObj *tail.Tail
)

//创建tailTask   一个日志收集任务
type TailTask struct {
	path       string
	topic      string
	instance   *tail.Tail
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	//根据路径去打开日志
	tailObj.init()
	return
}

func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed! err: ", err)
		return
	}
	go t.run()
}
func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task :%s_%s over!\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines:
			//3.1、循环每个日志收集项，发往kafka
			kafka.SendToChan(t.topic, line.Text) //函数调函数
		}

	}
}
