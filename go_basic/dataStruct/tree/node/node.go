package node

type Node struct {
	Data        string
	Left, right *Node
}

func NewNode(data string, left, right *Node) *Node {
	return &Node{
		Data:  data,
		Left:  left,
		right: right,
	}
}
