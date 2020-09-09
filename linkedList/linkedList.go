package linkedList

import "fmt"

//AddLast     O(n)
//AddFirst    O(1)
//Add         O(n)
//RemoveLast  O(n)
//RemoveFirst O(1)
//Remove      O(n)
//Set         O(n)
//Get         O(n)
//Contains    O(n)

type LinkedList struct {
	DummyHead *Node
	size      int
}

type Node struct {
	value interface{}
	next  *Node
}

func NewLNodeWithValueAndNext(value interface{}, next *Node) *Node {
	return &Node{
		value: value,
		next:  next,
	}
}

func NewNodeValue(value interface{}) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

func NewNode() *Node {
	return &Node{
		value: nil,
		next:  nil,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		DummyHead: NewNode(),
		size:      0,
	}
}

func (l *LinkedList) String() string {
	result := "linked list : ["
	if l.size == 0 {
		result += "empty!]"
		return result
	}

	cur := l.DummyHead.next
	for i := 0; i < l.size; i++ {
		result += fmt.Sprintf("%v->", cur.value)
		cur = cur.next
	}
	result += "null]"
	return result
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) AddFirst(value interface{}) {
	//node:=NewNode()
	//node.next = l.DummyHead
	//l.DummyHead = node

	_ = l.Add(value, 1)
}

func (l *LinkedList) Add(value interface{}, index int) error {
	if index-1 < 0 || index-1 > l.size {
		return invalidIndex
	}

	//if index == 0 {
	//	l.AddFirst(value)
	//	return nil
	//}

	var prev *Node = l.DummyHead
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}

	prev.next = NewLNodeWithValueAndNext(value, prev.next)
	l.size++

	return nil
}

func (l *LinkedList) AddLast(value interface{}) {
	_ = l.Add(value, l.size+1)
}

func (l *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index > l.size {
		return nil, invalidIndex
	}

	cur := l.DummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur.value, nil
}

func (l *LinkedList) GetFirst() (interface{}, error) {
	return l.Get(0)
}

func (l *LinkedList) GetLast() (interface{}, error) {
	return l.Get(l.size)
}

func (l *LinkedList) Set(value interface{}, index int) error {
	if index < 0 || index > l.size {
		return invalidIndex
	}

	cur := l.DummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	cur.value = value
	return nil
}

func (l *LinkedList) Contains(value interface{}) bool {

	cur := l.DummyHead.next
	for i := 0; i < l.size; i++ {
		if cur.value == value {
			return true
		}

		cur = cur.next
	}

	return false
}

func (l *LinkedList) Remove(index int) (interface{}, error) {
	if index-1 < 0 || index-1 > l.size {
		return nil, invalidIndex
	}

	prev := l.DummyHead
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}
	delNode := prev.next
	prev.next = delNode.next
	l.size--

	return delNode.value, nil
}

func (l *LinkedList) RemoveFirst() (interface{}, error) {
	return l.Remove(1)
}

func (l *LinkedList) RemoveLast() (interface{}, error) {
	return l.Remove(l.size)
}
