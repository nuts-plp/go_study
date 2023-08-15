package seqQueue

import "fmt"

type Node struct {
	Data any
}

type queue struct {
	num         int
	front, rear int
	list        []*Node
}

func NewQueue() *queue {
	q := &queue{
		0,
		0, 0,
		make([]*Node, 3, 3),
	}
	return q
}
func (q *queue) Add(D any) {
	p := &Node{
		D,
	}
	if q.front == q.rear {
		l := make([]*Node, cap(q.list)*2, cap(q.list)*2)
		j := 0
		for i := q.front; i < q.num+q.front; i++ {
			fmt.Println("-------------===", cap(l))
			l[j] = q.list[i]
		}
		q.list = l
		q.front = 0
		q.rear = q.num % cap(q.list)
	}
	fmt.Println(q.list)
	q.list[q.rear] = p
	q.num++
	q.rear = (q.rear + 1) % cap(q.list)
}
func (q *queue) Poll() any {
	if q.IsEmpty() {
		return nil
	}
	p := q.list[q.front].Data
	q.front = (q.front + 1) % cap(q.list)
	q.num--
	return p
}
func (q *queue) Peek() any {
	if q.IsEmpty() {
		return nil
	}
	return q.list[q.front].Data

}
func (q *queue) IsEmpty() bool {
	return q.num == 0
}
