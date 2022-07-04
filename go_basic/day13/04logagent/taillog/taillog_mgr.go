package taillog

import (
	"fmt"
	"go_basic/day13/04logagent/etcd"
	"time"
)

var taskMger *Manager

type Manager struct {
	logEntry    []*etcd.LogEntry
	taskMap     map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	taskMger = &Manager{
		logEntry:    logEntryConf,
		taskMap:     make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry),
	}
	for _, logEntry := range logEntryConf {
		taskObj := NewTailTask(logEntry.Path, logEntry.Topic)

		mKey := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		taskMger.taskMap[mKey] = taskObj
	}

	go taskMger.run()
}

//监听自己的newConfChan，有了新的配置之后做对应处理
func (t *Manager) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mKey := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.taskMap[mKey]
				if ok {
					//原来就有，不需要操作
					continue
				} else {
					//新增的
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.taskMap[mKey] = tailObj
				}
				for _, c1 := range t.logEntry {
					isDelete := true
					for _, c2 := range newConf {
						if c2.Path == c1.Path && c2.Topic == c1.Topic {
							isDelete = false
							continue
						}
					}
					if isDelete {
						//把c1的这个taskObj停掉
						mKey := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
						t.taskMap[mKey].cancelFunc()
					}
				}
			}
			//找出原来有但是现在 NewConf 没有的，要删掉

			//1、配置新增
			//2、配置删除
			//3、配置变更
			fmt.Println("新的配置来了！", newConf)
		default:
			time.Sleep(time.Second)

		}
	}
}

//NewConfChan 该函数向外暴露taskMger的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return taskMger.newConfChan
}
