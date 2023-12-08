package RBTree

// 定义红黑常量
const (
	RED   = true
	BLACK = false
)

//红黑树结构
type RBNode struct {
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
	Color  bool
	//DataItem interface{}
	Item //数据

}

//数据接口
type Item interface {
	Less(than Item) bool
}

//红黑树
type RBTree struct {
	NIL   *RBNode
	Root  *RBNode
	count uint
}

//比大小
func less(x, y Item) bool {
	return x.Less(y)
}

func NewRBTree() *RBTree {
	return new(RBTree).Init()
}

func (rbt *RBTree) Init() *RBTree {
	node := &RBNode{nil, nil, nil, BLACK, nil}
	return &RBTree{node, node, 0}
}

// 获取红黑树的长度
func (rbt *RBTree) Len() uint {
	return rbt.count
}

//取得红黑树的极大值
func (rbt *RBTree) max(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}
	for {
		if x.Right != rbt.NIL {
			x = x.Right
		}
	}
	return x
}

//取得红黑树的极小值
func (rbt *RBTree) min(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}
	for {
		if x.Left != rbt.NIL {
			x = x.Left
		}
	}
	return x
}

// 搜索红黑树
func (rbt *RBTree) search(x *RBNode) *RBNode {
	pnode := rbt.Root
	for {
		if less(pnode.Item, x.Item) {
			pnode = pnode.Right
		} else if less(x.Item, pnode.Item) {
			pnode = pnode.Left
		} else {
			break
		}
	}
	return pnode
}

func (rbt *RBTree) leftRotate(x *RBNode) {
	if x.Right == rbt.NIL {
		return //左旋，逆时针，右孩子不可以为零
	}
	y := x.Right
	x.Right = y.Left
}
