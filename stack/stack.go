package stack

type Stack interface {
	GetSize() int                   //O(1)
	IsEmpty() bool                  //O(1)
	Enqueue(interface{}) error      //O(1)
	Dequeue() (interface{}, error)  //O(1)
	GetFront() (interface{}, error) //O(1)
}
