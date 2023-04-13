package session

import (
	"errors"
	"sync"
)

//对象
//MemorySession设计
//定义MemorySession对象（字段名：sessionId，存kv的map，读写锁）
//构造函数，为了获取对象
//Set()
//Get()
//Del()
//Save()
type MemorySession struct {
	SessionId string
	data      map[string]interface{}
	rwlock    sync.RWMutex
}

//NewMemorySession 构造器
func NewMemorySession(sessionId string) *MemorySession {
	return &MemorySession{
		SessionId: sessionId,
		data:      make(map[string]interface{}, 16),
	}
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	//加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists!")
		return
	}
	return
}

func (m *MemorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return
}
func (m *MemorySession) Save() (err error) {
	return
}
