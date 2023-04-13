package AVLtree

import (
	"errors"
	"fmt"
)

/*
	AVL树适用于没有删除的场景
	红黑树的增删改查最优


*/
type AVLNode struct {
	Data  interface{} //数据
	Left  *AVLNode
	Right *AVLNode
	Heght int //高度
}

//comparator 函数指针类型
type comparator func(a, b interface{}) int

//compare函数指针
var compare comparator

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func NewNode(data interface{}) *AVLNode {
	return &AVLNode{
		data,
		nil,
		nil,
		1,
	}
}

func NewAVLTree(data interface{}, myfunc comparator) (*AVLNode, error) {
	if data == nil && myfunc == nil {
		return nil, errors.New("不能为空")
	}
	compare = myfunc
	return NewNode(data), nil
}

// 左旋
func (avlnode *AVLNode) LeftRotate() *AVLNode {
	var node *AVLNode

	return node
}

// 右旋
func (avlnode *AVLNode) RightRotate() *AVLNode {
	var node *AVLNode
	node = avlnode.Left
	avlnode.Left = node.Right
	node.Right = avlnode
	avlnode.Heght = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight()) + 1
	node.Heght = Max(node.Left.GetHeight(), node.Right.GetHeight()) + 1
	return node

}

// 先左旋再右旋
func (avlnode *AVLNode) LeftThenRightRotate() *AVLNode {
	var node *AVLNode
	node = avlnode.Left.LeftRotate()
	avlnode.Left = node
	return avlnode.RightRotate()
}

//先右旋再左旋
func (avlnode *AVLNode) RightThenLeftRotate() *AVLNode {
	var node *AVLNode
	node = avlnode.Right.RightRotate()
	avlnode.Right = node
	return avlnode.LeftRotate()
}

// 树的自动调整，自动平衡
func (avlnode *AVLNode) adjust() *AVLNode {
	if avlnode.Right.GetHeight()-avlnode.Left.GetHeight() == 2 {
		if avlnode.Right.Right.GetHeight() > avlnode.Right.Left.GetHeight() {
			avlnode = avlnode.LeftRotate()
		} else {
			avlnode = avlnode.RightThenLeftRotate()
		}
	} else if avlnode.Left.GetHeight()-avlnode.Right.GetHeight() == 2 {
		if avlnode.Left.Left.GetHeight() > avlnode.Left.Right.GetHeight() {
			avlnode = avlnode.RightRotate()
		} else {
			avlnode = avlnode.LeftThenRightRotate()
		}
	}
	return avlnode
}

// 数据的插入
func (avlnode *AVLNode) Insert(value interface{}) *AVLNode {
	if avlnode == nil {
		node := &AVLNode{value, nil, nil, 1}
		return node
	}
	switch compare(value, avlnode.Data) {
	case -1:
		avlnode.Left = avlnode.Left.Insert(value)
		avlnode = avlnode.adjust()
	case 1:
		avlnode.Right = avlnode.Right.Insert(value)
		avlnode = avlnode.adjust()
	case 0:
		fmt.Println("数据已存在")
	}
	avlnode.Heght = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight()) + 1
	return avlnode
}

// 数据的删除
func (avlnode *AVLNode) Delete(value interface{}) *AVLNode {
	if avlnode == nil {
		node := &AVLNode{value, nil, nil, 1}
		return node
	}
	switch compare(value, avlnode.Data) {
	case -1:
		avlnode.Left = avlnode.Left.Delete(value)
	case 1:
		avlnode.Right = avlnode.Right.Delete(value)
	case 0:
		//删除操作
		if avlnode.Left != nil && avlnode.Right != nil { //左右都右节点
			avlnode.Data = avlnode.Right.FindMin().Data
			avlnode.Right = avlnode.Right.Delete(avlnode.Data)
		} else if avlnode.Left != nil { //左孩子存在，右孩子存在或者不存在
			avlnode = avlnode.Left
		} else { //只有一个右孩子，或者无孩子
			avlnode = avlnode.Right
		}
	}
	return avlnode
}

func (avlnode *AVLNode) Find(data interface{}) *AVLNode {
	var node *AVLNode = nil
	switch compare(data, avlnode.Data) {
	case -1:
		node = avlnode.Left.Find(data)
	case 1:
		node = avlnode.Right.Find(data)
	case 0:
		node = avlnode
	}
	return node
}

//查找最小节点
func (avlnode *AVLNode) FindMin() *AVLNode {
	var node *AVLNode
	if avlnode.Left != nil {
		node = avlnode.Left.FindMin()
	} else {
		node = avlnode
	}
	return node
}

//查找最大节点
func (avlnode *AVLNode) FindMax() *AVLNode {
	var node *AVLNode
	if avlnode.Right != nil {
		node = avlnode.Right.FindMax()
	} else {
		node = avlnode
	}
	return node
}

//抓取数据
func (avlode *AVLNode) GetData() interface{} {
	return avlode.Data
}

//设置
func (avlnode *AVLNode) SetData(data interface{}) {
	avlnode.Data = data
}
func (avlnode *AVLNode) GetLeft() *AVLNode {
	if avlnode == nil {
		return nil
	}
	return avlnode.Left
}

func (avlnode *AVLNode) GetRight() *AVLNode {
	if avlnode == nil {
		return nil
	}
	return avlnode.Right
}
func (avlnode *AVLNode) GetHeight() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.Heght
}

// getall
func (avlnode *AVLNode) GetAll() []interface{} {
	values := make([]interface{}, 0)
	return values
}
func AddValues(values []interface{}, avlnode *AVLNode) []interface{} {
	if avlnode != nil {
		values = AddValues(values, avlnode.Left)
		values = append(values, avlnode.Data)
		values = AddValues(values, avlnode.Right)
	}
	return values
}
