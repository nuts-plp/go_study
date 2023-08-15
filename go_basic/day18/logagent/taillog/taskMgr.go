package taillog

import (
	"fmt"
	"go_basic/day18/logagent/etcd"
)

var (
	taskMger *taskMgr
)

type taskMgr struct {
	logEntries  []*etcd.LogEntry
	taskMap     map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logConfEntry []*etcd.LogEntry) {
	taskMger = &taskMgr{
		logEntries:  logConfEntry,
		taskMap:     make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry),
	}

	for _, v := range logConfEntry {
		taskObj := NewTailTask(v.Path, v.Topic)
		taskMger.taskMap[fmt.Sprintf("%s+%s", v.Path, v.Topic)] = taskObj
	}
	go taskMger.run()
}

//GetNewConfChan 向外暴露newConfChan
func GetNewConfChan() chan<- []*etcd.LogEntry {
	return taskMger.newConfChan
}
func (t *taskMgr) run() {
	select {
	case newConf := <-t.newConfChan:
		for _, c := range newConf {
			mk := fmt.Sprintf("%s+%s", c.Path, c.Topic)
			_, ok := t.taskMap[mk]
			if ok {
				//任务中本来就有
				continue
			} else {
				taskObj := NewTailTask(c.Path, c.Topic)
				t.taskMap[mk] = taskObj
			}
		}
		for _, v1 := range t.logEntries {
			for _, v3 := range newConf {
				mk := fmt.Sprintf("%s+%s", v1.Path, v1.Topic)
				if v1.Topic != v3.Topic && v1.Path != v3.Path {
					t.taskMap[mk].Cancel()
				}
			}
		}
	}
}
