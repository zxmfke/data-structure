package queue

import "fmt"

/**
 * @Author: Zheng xiaomin
 * @Date: 2020/9/8 20:00 晚上
 */

// GetSize() int                   O(1)
// IsEmpty() bool                  O(1)
// Enqueue(interface{}) error      O(1)
// Dequeue() (interface{}, error)  O(n)  第一个值拿出来后，后面都要往前移一位
// GetFront() (interface{}, error) O(1)

type ArrayQueue struct {
	data        []interface{}
	size        int
	capacity    int
	allowResize bool
}

func NewArrayQueue(capacity int) *ArrayQueue {
	data := make([]interface{}, 0, capacity)
	return &ArrayQueue{
		data:        data,
		size:        0,
		capacity:    capacity,
		allowResize: false,
	}
}

func NewDefaultArrayQueue() *ArrayQueue {
	data := make([]interface{}, 0, 5)
	return &ArrayQueue{
		data:        data,
		size:        0,
		capacity:    5,
		allowResize: false,
	}
}

func (a *ArrayQueue) String() string {
	return fmt.Sprintf("\n current stack size is %d, capacity is : %d \n stack : %v  tail", a.size, a.capacity, a.data)
}

func (a *ArrayQueue) WithData(arr []int) *ArrayQueue {
	if a.size != 0 {
		return a
	}

	for i := 0; i < a.capacity; i++ {
		_ = a.addLast(arr[i])
	}

	return a
}

func (a *ArrayQueue) Length() int {
	return len(a.data)
}

func (a *ArrayQueue) GetSize() int {
	return a.size
}

func (a *ArrayQueue) Capacity() int {
	return a.capacity
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.size == 0
}

func (a *ArrayQueue) Enqueue(value interface{}) error {
	return a.addLast(value)
}

func (a *ArrayQueue) Dequeue() (interface{}, error) {
	return a.removeFirst()
}

func (a *ArrayQueue) GetFront() (interface{}, error) {
	return a.get(0)
}

func (a *ArrayQueue) addLast(value interface{}) error {
	if err := a.add(value, a.Length()+1); err != nil {
		return queueIsFull
	}

	return nil
}

func (a *ArrayQueue) add(value interface{}, index int) error {

	if index < 0 || index-2 > a.capacity {
		return invalidIndex
	}

	if a.size == a.Capacity() {
		if !a.allowResize {
			return queueIsFull
		} else {
			a.capacity *= 2
		}
	}

	a.data = append(a.data, 0)

	for i := a.size - 1; i >= index-1; i-- {
		if i == -1 {
			break
		}

		a.data[i+1] = a.data[i]
	}

	a.data[index-1] = value
	a.size++

	return nil

}

func (a *ArrayQueue) removeFirst() (interface{}, error) {
	v, err := a.remove(1)
	if err != nil {
		return 0, queueIsEmpty
	}

	return v, nil
}

func (a *ArrayQueue) remove(index int) (interface{}, error) {
	if index < 0 || index > a.size {
		return 0, invalidIndex
	}

	if a.size == 0 {
		return 0, queueIsEmpty
	}

	ret := a.data[index-1]

	for i := index; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}

	a.size--

	a.data = a.data[:a.size]

	// just prevent add last and remove last frequently when size +1 = capacity
	if a.allowResize {
		if a.size <= (a.capacity/4) && a.size/2 != 0 {
			a.resize()
		}
	}

	return ret, nil

}

func (a *ArrayQueue) resize() {
	tmpData := make([]interface{}, a.size, a.capacity/2)
	copy(tmpData, a.data[:a.size])
	a.data = tmpData
	a.capacity /= 2
}

func (a *ArrayQueue) get(index int) (interface{}, error) {
	if index < 0 || index > a.size {
		return -1, invalidIndex
	}

	return a.data[index-1], nil
}
