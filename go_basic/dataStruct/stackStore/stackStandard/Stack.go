package stack

type Stack interface {
	IsEmpty() bool
	//NewStack() any
	Pop() any
	Peek() any
	Push(D any)
}
