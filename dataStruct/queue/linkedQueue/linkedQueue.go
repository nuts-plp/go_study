package linkedQueue

type Node struct {
	Data any
	Next *Node
}
type queue struct {
	num  int
	head *Node
	tail *Node
}

func NewQueue() *queue {
	p := &Node{
		nil,
		nil,
	}
	return &queue{
		0,
		p,
		p,
	}
}
func (q *queue) IsEmpty() bool {
	return q.num == 0
}

func (q *queue) Poll() any {
	if q.IsEmpty() {
		return nil
	}
	p := q.head.Data
	q.head = q.head.Next
	return p
}
func (q *queue) Peek() any {
	if q.IsEmpty() {
		return nil
	}
	return q.head.Data
}
func (q *queue) Add(D any) {
	if q.IsEmpty() {
		q.tail.Data = D
		q.num++
		return
	}
	p := &Node{
		D,
		nil,
	}
	q.tail.Next = p
	q.tail = q.tail.Next
	q.num++
}
