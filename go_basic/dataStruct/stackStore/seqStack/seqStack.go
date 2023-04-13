package seqStack

type Node struct {
	Data any
}

type SeqStack struct {
	list []*Node
	num  int
}

func NewStack() *SeqStack {
	return &SeqStack{
		make([]*Node, 20, 20),
		0,
	}
}
func (s *SeqStack) IsEmpty() bool {
	return s.num == 0
}
func (s *SeqStack) Peek() any {
	if s.num == 0 {
		return nil
	}
	return s.list[s.num-1].Data
}
func (s *SeqStack) Pop() any {
	if s.num == 0 {
		return nil
	}
	s.num -= 1
	p := s.list[s.num]
	s.list = s.list[:s.num]
	return p.Data
}

// Push 向数组添加元素，如果数组已满则以二倍进行扩容
func (s *SeqStack) Push(D any) {
	if s.num == len(s.list) {
		p := make([]*Node, len(s.list)*2, len(s.list)*2)
		for i, v := range s.list {
			p[i] = v
		}
		s.list = p
	}
	n := &Node{
		D,
	}
	s.list[s.num] = n
	s.num += 1
}
