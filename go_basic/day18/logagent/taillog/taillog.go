package taillog

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"go_basic/day18/logagent/kafka"
	"time"
)

type TailTask struct {
	Path    string
	Topic   string
	Tail    *tail.Tail
	Context context.Context
	Cancel  context.CancelFunc
}

//NewTailTask 创建一个任务,并创建一个goroutine去不断的读取日志，返回一个任务对象
func NewTailTask(path, topic string) (taskObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	taskObj = &TailTask{
		Path:    path,
		Topic:   topic,
		Context: ctx,
		Cancel:  cancel,
	}
	taskObj.Init()
	go taskObj.run()
	mk := fmt.Sprintf("%s+%s", taskObj.Path, taskObj.Topic)
	taskMger.taskMap[mk] = taskObj
	return
}

func (t *TailTask) Init() {
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		ReOpen:    true,
		Follow:    true,
		Poll:      true,
	}
	tail, err := tail.TailFile(t.Path, config)
	if err != nil {
		fmt.Println("[taillog] init tail failed")
		return
	}
	t.Tail = tail
}
func (t *TailTask) run() {
	fmt.Println("new tailTask run...")
	for {
		select {
		case <-t.Context.Done():
			fmt.Printf("[taillog] tailTask:%s_%s finished\n", t.Path, t.Topic)
			return
		case line := <-t.Tail.Lines:
			fmt.Println(line.Text)
			kafka.SendToChan(t.Topic, line.Text)
		default:
			time.Sleep(time.Millisecond * 50)

		}
	}
}
