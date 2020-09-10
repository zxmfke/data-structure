package stack

import "fmt"


//Push O(1)
//Pop  O(1)
//Peek O(1)

type LinkedListStack struct {
	list *LinkedList
	size int
}

type LinkedList struct {
	DummyHead *Node
}

type Node struct {
	value interface{}
	next  *Node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		list: &LinkedList{
			DummyHead: NewNode(),
		},
		size: 0,
	}
}

func NewNodeWithValueAndNext(value interface{}, next *Node) *Node {
	return &Node{
		value: value,
		next:  next,
	}
}

func NewNode() *Node {
	return &Node{
		value: nil,
		next:  nil,
	}
}

func NewNodeValue(value interface{}) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

func (l *LinkedListStack) String() string {
	result := "linked list : ["
	if l.size == 0 {
		result += "empty!]"
		return result
	}

	cur := l.list.DummyHead.next
	for i := 0; i < l.size; i++ {
		result += fmt.Sprintf("%v->", cur.value)
		cur = cur.next
	}
	result += "null]"
	return result
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedListStack) GetSize() int {
	return l.size
}

func (l *LinkedListStack) Push(value interface{}) error {
	//node:=NewNode()
	//node.next = l.DummyHead
	//l.DummyHead = node

	return l.add(value, 1)
}

func (l *LinkedListStack) add(value interface{}, index int) error {
	if index-1 < 0 || index-1 > l.size {
		return invalidIndex
	}

	//if index == 0 {
	//	l.AddFirst(value)
	//	return nil
	//}

	var prev *Node = l.list.DummyHead
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}

	prev.next = NewNodeWithValueAndNext(value, prev.next)
	l.size++

	return nil
}

func (l *LinkedListStack) Pop() (interface{}, error) {
	//node:=NewNode()
	//node.next = l.DummyHead
	//l.DummyHead = node

	return l.remove(l.size)
}

func (l *LinkedListStack) remove(index int) (interface{}, error) {
	if index-1 < 0 || index-1 > l.size {
		return nil, invalidIndex
	}

	prev := l.list.DummyHead
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}
	delNode := prev.next
	prev.next = delNode.next
	l.size--

	return delNode.value, nil
}

func (l *LinkedListStack) get(index int) (interface{}, error) {
	if index < 0 || index > l.size {
		return nil, invalidIndex
	}

	cur := l.list.DummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur.value, nil
}

func (l *LinkedListStack) Peek() (interface{}, error) {
	return l.get(0)
}
