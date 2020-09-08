package queue

import "fmt"

/**
 * @Author: Zheng xiaomin
 * @Date: 2020/9/8 21:10 晚上
 */

// GetSize() int                   O(1)
// IsEmpty() bool                  O(1)
// Enqueue(interface{}) error      O(1)
// Dequeue() (interface{}, error)  O(1)
// GetFront() (interface{}, error) O(1)

type LoopQueue struct {
	data        []interface{}
	front       int
	tail        int
	size        int
	capacity    int
	allowResize bool
}

func NewLoopQueue(capacity int) *LoopQueue {
	data := make([]interface{}, 0, capacity+1)
	return &LoopQueue{
		data:        data,
		front:       0,
		tail:        0,
		size:        0,
		capacity:    capacity + 1,
		allowResize: false,
	}
}

func (l *LoopQueue) String() string {
	res := "["
	for i := 0; i < l.tail; i = (i + 1) % l.Capacity() {
		res = fmt.Sprintf("%s%s", res, l.data[i])
		if i+1 != l.front {
			res = fmt.Sprintf("%s,", res)
		}
	}
	res += "]"
	return fmt.Sprintf("\n current queue size is %d, capacity is : %d \n queue : %v  tail", l.size, l.capacity, res)
}

func (l *LoopQueue) Capacity() int {
	return l.capacity
}

func (l *LoopQueue) IsEmpty() bool {
	return l.front == l.tail
}

func (l *LoopQueue) IsFull() bool {
	if l.IsEmpty() {
		return false
	}

	return (l.tail+1)%l.Capacity() == l.front
}

func (l *LoopQueue) WithData(arr []int) *LoopQueue {
	if l.size != 0 {
		return l
	}

	for i := 0; i < l.capacity; i++ {
		_ = l.Enqueue(arr[i])
	}

	return l
}

func (l *LoopQueue) WithResize() *LoopQueue {
	l.allowResize = true
	return l
}

func (l *LoopQueue) GetSize() int {
	return l.size
}

func (l *LoopQueue) Enqueue(value interface{}) error {
	if l.IsFull() {
		if !l.allowResize {
			return queueIsFull
		} else {
			l.resize(l.capacity * 2)
		}
	}

	l.data = append(l.data, 0)
	l.data[l.tail] = value
	l.tail = (l.tail + 1) % l.Capacity()
	l.size++

	return nil
}

func (l *LoopQueue) resize(newCapacity int) {
	newArray := make([]interface{}, 0, newCapacity)
	for i := 0; i < l.size; i++ {
		newArray = append(newArray, l.data[(i+l.front)%l.Capacity()])
	}

	l.data = newArray
	l.front = 0
	l.tail = l.size
}

func (l *LoopQueue) Dequeue() (interface{}, error) {
	if l.IsEmpty() {
		return nil, queueIsEmpty
	}

	ret := l.data[l.front]
	l.front = (l.front + 1) % l.Capacity()
	l.size++

	if l.allowResize {
		if (l.size == l.Capacity()/4) && l.Capacity()/2 != 0 {
			l.resize(l.Capacity() / 2)
		}
	}

	return ret, nil
}

func (l *LoopQueue) GetFront() (interface{}, error) {
	if l.IsEmpty() {
		return nil, queueIsEmpty
	}

	return l.data[l.front], nil
}
