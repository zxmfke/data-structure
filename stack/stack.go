package stack

type Stack interface {
	GetSize() int               //O(1)
	IsEmpty() bool              //O(1)
	Push(interface{}) error     //O(1)
	Pop() (interface{}, error)  //O(1)
	Peek() (interface{}, error) //O(1)
}
