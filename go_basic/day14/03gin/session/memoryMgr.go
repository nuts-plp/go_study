package session

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

type MemorySessionMgr struct {
	sessionMap map[string]interface{}
	rwlock     sync.RWMutex
}

//NewMemorySessionMgr 构造函数
func NewMemorySessionMgr() *MemorySessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]interface{}, 1024),
	}
}

func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}
func (m *MemorySessionMgr) CreateSession(session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	//go get github.com/satori/go.uuid
	id := uuid.NewV4()
	sessionId := id.String()
	session = NewMemorySession(sessionId)

}
func (m *MemorySessionMgr) Get(sessionId string, err error) {
	return
}
