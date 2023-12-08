package seqListMgr_slice

import (
	"errors"
	"fmt"
	per "go_basic/dataStruct/list/standard"
)

type Manager struct {
	person []*per.Person
}

var (
	seqlist *Manager
	//元素数量
	size int
	//容量
	capitaty int
	//切片是否可扩展
	expandable bool
)

//NewSeqListMgr_slice 构造器 初始化管理者 为切片开辟空间
func NewSeqListMgr_slice(cap int, expand bool) *Manager {
	seqlist = &Manager{
		person: make([]*per.Person, 5, 10),
	}
	capitaty = cap
	expandable = expand
	return seqlist
}
func (s *Manager) Add(data string) (err error) {
	err = nil
	if capitaty > size { //容量还足
		s.person[size] = per.NewPerson(data)
		size++
		return err
	} else {            //容量不足
		if expandable { //可扩展
			_ = append(seqlist.person, &per.Person{
				Data: data,
			})
			return err
		} else {
			err = errors.New("expand seqList failed")
			return err
		}
	}

}

//Search 查找对应顺序表索引的数据
func (s *Manager) Search(index int) (data interface{}, err error) {
	err = nil
	//判断顺序表中是否有元素
	if size == 0 {
		err = errors.New("data not exists")
		return nil, err
	}
	//判断传入的索引是否越界
	if index >= size || index <= 0 {
		err = errors.New("index out of bound")
		return nil, err
	}
	return s.person[index].Data, err
}

//Delete 删除对应索引的结点
func (s *Manager) Delete(index int) (err error) {
	err = nil
	if size == 0 {
		err = errors.New("seqList is null")
		return
	}
	if index < 0 || index >= size {
		err = errors.New("index out of bound")
		return
	}
	//for 循环，删除目标结点后把目标节点后的结点前移
	for i := index; i < size; i++ {
		// 退出循环
		if i == size-1 {
			s.person[i].Data = ""
			break
		}
		s.person[i] = s.person[i+1]
	}
	size--
	//把最后一个结点置空

	return nil
}

//Modify 修改索引结点的数据
func (s *Manager) Modify(index int, data string) (err error) {
	err = nil
	if size == 0 {
		err = errors.New("seqList is null")
		return err
	}

	if index < 0 || index >= size {
		err = errors.New("index out of bound")
		return err
	}
	s.person[index].Data = data
	return err
}

func (s *Manager) ToString() {

	for i, v := range s.person {
		fmt.Printf("[index:%d、data:%v]", i, v.Data)
	}
	fmt.Println()

}
func (s *Manager) GetSize() int {
	return size
}

////isNull 判断列表是否为空列表
//func isNull()bool{
//	return capitaty==0
//}
