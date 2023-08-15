package queue

type Queue interface {
	Add(Data any)
	Poll() any
	Peek() any
	IsEmpty() bool
}
