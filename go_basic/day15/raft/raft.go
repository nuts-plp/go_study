package main

import (
	"sync"
)

//实现3结点选举
//RPC 分布式实现
//完整代码 自主选举 日志复制

type Raft struct {
	mu              sync.Mutex
	me              int
	currentTerm     int
	voteFor         int
	state           int
	lastMessageTime int64
	currentLeader   int
	message         chan bool
	electCh         chan bool
	heartBeat       chan bool
	heartBeatRe     chan bool
	timeout         int
}
