package seqListMgr_map

import (
	"errors"
	"fmt"
	per "go_basic/dataStruct/list/standard"
)

type Manager struct {
	person map[string]*per.Person
}

var (
	size    int
	capa    int
	manager *Manager
)

//NewSeqList_map 构造器
func NewSeqList_map(cap int) *Manager {
	manager = &Manager{
		person: make(map[string]*per.Person, cap),
	}
	capa = cap
	return manager
}

//Add 添加结点
func (s *Manager) Add(index, data string) (err error) {
	err = nil
	//map容量为空
	if capa == 0 {
		err = errors.New("invalid map")
		return
	}
	if capa == size {
		err = errors.New("map is full")
		return
	}
	//添加数据
	s.person[index] = per.NewPerson(data)
	//元素个数加1
	size++
	return
}

//Search 查找结点
func (s *Manager) Search(index string) (data interface{}, err error) {
	err = nil
	for i, v := range s.person {
		if i == index {
			return v, err
		}
	}
	err = errors.New("not found")
	return nil, err
}

//Delete 删除结点
func (s *Manager) Delete(index string) (err error) {
	err = nil
	for i := range s.person {
		if i == index {
			delete(s.person, index)
			size--
			return
		}
	}
	err = errors.New("without the index of node")
	return

}

//Modify 修改结点内容
func (s *Manager) Modify(index string, data string) (err error) {
	err = nil
	for i := range s.person {
		if i == index {
			s.person[index].Data = data
			return err
		}
	}
	err = errors.New("node of index not found")
	return
}

//GetSize 获取结点个数
func (s *Manager) GetSize() int {
	return size
}
func (s *Manager) ToString() {
	var str string
	for i, v := range s.person {
		str += fmt.Sprintf("[key:%v、value:%v]、", i, v.Data)
	}
	fmt.Println(str)
}
