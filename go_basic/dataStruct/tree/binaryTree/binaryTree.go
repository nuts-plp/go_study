package binaryTree

import "fmt"

type Node struct {
	Data  int
	Left  *Node //左节点
	Right *Node //右节点
	Next  *Node //相同节点
}

type BinaryTree struct {
	Root *Node
	Size int
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		&Node{Data: -1},
		0,
	}
}

// Add 向当前二叉树加入新的节点
func (bt *BinaryTree) Add(data int) {
	if bt.Size == 0 {
		bt.Root.Data = data
		bt.Size++
		return
	}
	node := &Node{
		Data: data,
	}
	bt.Root.add(node)
	bt.Size++
}

//添加节点，三种情况  1、比当前节点大  2、比当前节点小    3、和当前节点等大
func (n *Node) add(node *Node) *Node {
	//加入的节点比当前节点大
	if n.Data < node.Data {
		if n.Right == nil {
			n.Right = node
			return nil
		} else {
			n.Right.add(node)
		}
	}
	//加入的节点比当前节点数据小
	if n.Data > node.Data {
		if n.Left == nil {
			n.Left = node
			return nil
		} else {
			n.Left.add(node)
		}
	}
	//加入节点数据和当前节点数据大小相等
	if n.Data == node.Data {
		for {
			if n.Next == nil {
				n.Next = node
				return nil
			} else {
				n.Next = n.Next.Next
			}
		}
	}
	return nil
}

//判断二叉树是否为空
func (bt *BinaryTree) IsEmpty() bool {
	return bt.Size == 0
}

//返回数据所在节点的层次
func (bt *BinaryTree) GetLevel(data int) int {
	return 0
}

//返回树的高度
func (bt *BinaryTree) GetHeight() int {
	return 0
}

func (n *Node) preOrder() {
	fmt.Println(n.Data)
	if n.Left != nil {
		n.Left.preOrder()
	}
	if n.Right != nil {
		n.Right.preOrder()
	}
	return
}

//二叉树的先根次序遍历   递归
func (bt *BinaryTree) PreOrder() {
	if bt.Size != 0 {
		bt.Root.preOrder()
	}
	return
}
func (n *Node) inOrder() {
	if n.Left != nil {
		n.Left.preOrder()
	}
	fmt.Println(n.Data)
	if n.Right != nil {
		n.Right.preOrder()
	}
	return
}

//树的中根次序便利  递归
func (bt *BinaryTree) InOrder() {
	if bt.Size != 0 {
		bt.Root.inOrder()
	}
	return
}
func (n *Node) postOrder() {
	if n.Left != nil {
		n.Left.preOrder()
	}
	if n.Right != nil {
		n.Right.preOrder()
	}
	fmt.Println(n.Data)
	return
}

//二叉树的后根遍历次序 递归
func (bt *BinaryTree) PostOrder() {
	if bt.Size != 0 {
		bt.Root.postOrder()
	}
	return
}

//二叉树的层次遍历  把各个节点存入队列进行遍历，借助队列先入先出的特点
func (bt *BinaryTree) LevelOrder() {}

//树的先根、中根、后根次序遍历，非递归

//树的节点的子树的删除
func (bt *BinaryTree) Remove(node *Node) {
	if node != nil && bt.Contain(node) {
		node.Left = nil
		node.Right = nil
	}
	return
}

//递归删除树
func (bt *BinaryTree) Clear() {
	if bt.Size != 0 {
		bt.Root.Right = nil
		bt.Root.Left = nil
		bt.Root.Data = -1
		bt.Size = 0
		return
	}

}

//判断二叉树是否包含某个节点
func (bt *BinaryTree) Contain(node *Node) bool {
	return false
}
