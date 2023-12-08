package linkStack

type any = interface{}
type Node struct {
	Data any
	Pre  *Node
}
type LinkStack struct {
	Pointer *Node
}

func newNode() *Node {
	return &Node{}
}
func NewStack() *LinkStack {
	p := LinkStack{
		newNode(),
	}
	return &p
}

// Pop 取栈顶元素
func (s *LinkStack) Pop() any {
	n := s.Pointer
	s.Pointer = s.Pointer.Pre
	return n.Data
}

// Peek 返回栈顶元素
func (s *LinkStack) Peek() any {
	return s.Pointer.Data
}

// Push 元素进栈
func (s *LinkStack) Push(D any) {
	if s.Pointer.Data == nil {
		s.Pointer.Data = D
		return
	}
	newNode := &Node{
		D,
		s.Pointer,
	}
	s.Pointer = newNode

}
func (s *LinkStack) IsEmpty() bool {

	return s.Pointer.Data == nil
}
