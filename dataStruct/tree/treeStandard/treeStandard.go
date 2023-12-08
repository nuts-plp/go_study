package treeStandard

type TreeStand interface {
	//判空
	IsEmpty() bool
	//返回根节点所在的层次
	level(data interface{}) int
	//返回树的结点数
	Size() int
	//返回树的高度
	Height() int
	//输出树的先根次序遍历
	preorder()
	//输出树的中根次序遍历
	inorder()
	//输出树的后根次序遍历
	postorder()
	//输出树的层次遍历序列
	levelorder()
	//插入x作为根节点
	insert(x interface{}) error
	//插入x作为p结点的第i个孩子
	insertS(x, p interface{}, i int) error
	//删除p结点的第i个子树
	remove(p interface{}, i int) error
	//删除树的所有结点
	clear()
	//查找首个关键字为key的结点
	search(key interface{}) error
	//判断是否包含关键字为key的结点
	contains(key interface{}) bool
	//删除以p结点为根结点的子树
	delete(p interface{}) error
}
