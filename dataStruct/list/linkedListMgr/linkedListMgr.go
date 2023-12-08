package linkedListMgr

import (
	"errors"
	"fmt"
	"go_basic/dataStruct/list/standard"
)

type Manager struct {
	person *standard.Person
}

var (
	manager *Manager
	size    int
	head    *standard.Person
)

func NewLinkedListMgr() *Manager {
	head = standard.NewPerson("")
	return manager
}
func (m *Manager) Add(data string) (err error) {
	pointer := head
	err = nil
	if size == 0 {
		head.Data = data
		size++
		return
	}
	for { //遍历下一个
		if pointer.Next == nil {
			pointer.Next = standard.NewPerson(data)
			size++
			return
		}
		pointer = pointer.Next
	}

}
func (m *Manager) Search(index int) (data interface{}, err error) {
	pointer := head
	err = nil
	//声明一个变量 指向当前遍历到的结点
	var n = 1
	if index > 0 && index <= size {
		for {
			if n == index {
				return pointer.Data, err
			}
			pointer = pointer.Next
			n++
		}
	}
	err = errors.New("without the node of index")
	return nil, err
}
func (m *Manager) Delete(index int) (err error) {
	err = nil
	pointer := head
	var n = 1
	//index在元素索引范围内
	if index > 0 && index <= size {
		//index==1
		if index == 1 {
			head = pointer.Next
			return
		}
		for {
			// 1<index<size
			if n+1 == index {
				pointer.Next = pointer.Next.Next
				return nil
			}
			//index==size
			if index == size {
				pointer.Next = nil
				return
			}
			pointer = pointer.Next
			n++
		}
	}
	//index在元素索引范围外
	err = errors.New("without the index of node")
	return

}
func (m *Manager) Modify(index int, data string) (err error) {
	pointer := head
	err = nil
	var n = 1
	//	1<=index<=size
	if index > 0 && index <= size {
		for {
			if n == index {
				pointer.Data = data
				return
			}
			n++
			pointer = pointer.Next
		}
	}
	err = errors.New("without the index of node")
	return
}
func (m *Manager) GetSize() int {
	return size
}
func (m *Manager) ToArray() {
	pointer := head
	var n = 1
	var str string
	for {
		if pointer.Next == nil {
			str += fmt.Sprintf("[key:%d、value:%v]、", n, pointer.Data)
			fmt.Println(str)
			return
		}
		str += fmt.Sprintf("[key:%d、value:%v]、", n, pointer.Data)
		pointer = pointer.Next
		n++
	}
}
