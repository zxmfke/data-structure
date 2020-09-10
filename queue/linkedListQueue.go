package queue

import "fmt"

/**
 * @Author: Zheng xiaomin
 * @Date: 2020/9/10 23:00 晚上
 */

// Enqueue(interface{}) error      O(1)
// Dequeue() (interface{}, error)  O(1)
// GetFront() (interface{}, error) O(1)

type LinkedListQueue struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	value interface{}
	next  *Node
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head: NewNode(),
		tail: nil,
		size: 0,
	}
}

func NewNodeWithValue(value interface{}) *Node {
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

func (l *LinkedListQueue) String() string {
	result := "linked list queue : front ["
	if l.size == 0 {
		result += "empty!]"
		return result
	}

	cur := l.head
	for i := 0; i < l.size; i++ {
		result += fmt.Sprintf("%v->", cur.value)
		cur = cur.next
	}
	result += "null]  tail"
	return result
}

func (l *LinkedListQueue) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedListQueue) GetSize() int {
	return l.size
}

func (l *LinkedListQueue) Enqueue(value interface{}) error {
	if l.tail == nil {
		l.tail = NewNodeWithValue(value)
		l.head = l.tail
	} else {
		l.tail.next = NewNodeWithValue(value)
		l.tail = l.tail.next
	}

	l.size++
	return nil
}

func (l *LinkedListQueue) Dequeue() (interface{}, error) {
	if l.IsEmpty() {
		return nil, queueIsEmpty
	}

	node := l.head
	l.head = l.head.next
	node.next = nil
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return node.value, nil
}

func (l *LinkedListQueue) GetFront() (interface{}, error) {
	if l.IsEmpty() {
		return nil, queueIsEmpty
	}

	return l.head.value, nil
}
