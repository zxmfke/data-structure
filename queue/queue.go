package queue

type Queue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
	GetFront() (interface{}, error)
}
