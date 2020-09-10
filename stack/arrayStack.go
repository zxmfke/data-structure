package stack

import "fmt"

/**
 * @Author: Zheng xiaomin
 * @Date: 2020/9/8 20:00 晚上
 */

//Push O(1)
//Pop  O(n)
//Peek O(1)

type ArrayStack struct {
	data        []interface{}
	size        int
	capacity    int
	allowResize bool
}

func NewArrayStack(capacity int) *ArrayStack {
	data := make([]interface{}, 0, capacity)
	return &ArrayStack{
		data:        data,
		size:        0,
		capacity:    capacity,
		allowResize: false,
	}
}

func NewDefaultArrayStack() *ArrayStack {
	data := make([]interface{}, 0, 5)
	return &ArrayStack{
		data:        data,
		size:        0,
		capacity:    5,
		allowResize: false,
	}
}

func (a *ArrayStack) String() string {
	return fmt.Sprintf("\n current stack size is %d, capacity is : %d \n stack : %v  TOP", a.size, a.capacity, a.data)
}

func (a *ArrayStack) WithData(arr []int) *ArrayStack {
	if a.size != 0 {
		return a
	}

	for i := 0; i < a.capacity; i++ {
		_ = a.addLast(arr[i])
	}

	return a
}

func (a *ArrayStack) Length() int {
	return len(a.data)
}

func (a *ArrayStack) GetSize() int {
	return a.size
}

func (a *ArrayStack) Capacity() int {
	return a.capacity
}

func (a *ArrayStack) IsEmpty() bool {
	return a.size == 0
}

func (a *ArrayStack) Push(value interface{}) error {
	return a.addLast(value)
}

func (a *ArrayStack) Pop() (interface{}, error) {
	return a.removeLast()
}

func (a *ArrayStack) Peek() (interface{}, error) {
	last := a.size - 1
	if last != -1 {
		return a.get(last)
	}

	return nil, stackIsEmpty
}

func (a *ArrayStack) addLast(value interface{}) error {
	if err := a.add(value, a.Length()+1); err != nil {
		return stackIsFull
	}

	return nil
}

func (a *ArrayStack) add(value interface{}, index int) error {

	if index < 0 || index-2 > a.capacity {
		return invalidIndex
	}

	if a.size == a.Capacity() {
		if !a.allowResize {
			return stackIsFull
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

func (a *ArrayStack) removeLast() (interface{}, error) {
	v, err := a.remove(a.size)
	if err != nil {
		return 0, stackIsEmpty
	}

	return v, nil
}

func (a *ArrayStack) remove(index int) (interface{}, error) {
	if index < 0 || index > a.size {
		return 0, invalidIndex
	}

	if a.size == 0 {
		return 0, stackIsEmpty
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

func (a *ArrayStack) resize() {
	tmpData := make([]interface{}, a.size, a.capacity/2)
	copy(tmpData, a.data[:a.size])
	a.data = tmpData
	a.capacity /= 2
}

func (a *ArrayStack) get(index int) (interface{}, error) {
	if index < 0 || index > a.size {
		return -1, invalidIndex
	}

	return a.data[index-1], nil
}
