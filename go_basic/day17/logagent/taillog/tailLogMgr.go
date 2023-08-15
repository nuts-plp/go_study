package taillog

import (
	"fmt"
	"go_basic/day17/logagent/etcd"
	"time"
)

var (
	taskMgr *tailLogMgr
)

type tailLogMgr struct {
	logEntries  []*etcd.LogEntry
	taskMap     map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntries []*etcd.LogEntry) {
	taskMgr = &tailLogMgr{
		logEntries:  logEntries,
		taskMap:     make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry),
	}
	for _, logEntry := range logEntries {
		//把启动的tailTask存入到map中
		tailTask := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		taskMgr.taskMap[mk] = tailTask
	}
	go taskMgr.run()
}

//监听自己的newConfChan，有了新的配置就做对应的处理
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.taskMap[mk]
				if ok {
					//原本就有
					continue
				} else {
					//有新增
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.taskMap[mk] = tailObj
				}
			}
			//找出原来logEntry有，但现在newConf没有的
			for _, c1 := range t.logEntries {
				isDelete := true
				for _, c2 := range newConf {
					if c1.Path == c2.Path && c1.Topic == c2.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.taskMap[mk].cancelFunc()
				}
			}
			//1 配置新增
			//2 配置删除
			//3 配置变更
			fmt.Println("新的配置来了", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

//
func NewConfChan() chan<- []*etcd.LogEntry {
	return taskMgr.newConfChan
}
