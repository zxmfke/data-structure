package arrayList

import "fmt"

/**
 * @Author: Zheng xiaomin
 * @Date: 2020/9/8 0:30 凌晨
 */

// 无resize的数组
//增 ： add/addFirst/addLast O(n/2)/O(n)/O(1) O(n)
//删 ： remove/removeFirst/removeLast O(n/2)/O(n)/O(1) O(n)
//改 ： set(value,index) O(1)
//查 ： get(index) O(1)
//find(value) O(n)
//contains(value) O(n)

type ArrayList struct {
	data        []int
	size        int
	capacity    int
	allowResize bool
}

func NewArrayList(capacity int) *ArrayList {
	data := make([]int, 0, capacity)
	return &ArrayList{
		data:        data,
		size:        0,
		capacity:    capacity,
		allowResize: false,
	}
}

func NewDefaultArrayList() *ArrayList {
	data := make([]int, 0, 5)
	return &ArrayList{
		data:        data,
		size:        0,
		capacity:    5,
		allowResize: false,
	}
}

func (a *ArrayList) String() string {
	return fmt.Sprintf("\n current array lis size is %d, capacity is : %d \n data : %v", a.size, a.capacity, a.data)
}

func (a *ArrayList) Length() int {
	return len(a.data)
}

func (a *ArrayList) Capacity() int {
	return a.capacity
}

func (a *ArrayList) WithData(arr []int) *ArrayList {
	if a.size != 0 {
		return a
	}

	for i := 0; i < a.capacity; i++ {
		_ = a.AddLast(arr[i])
	}

	return a
}

func (a *ArrayList) WithResize() *ArrayList {
	a.allowResize = true
	return a
}

func (a *ArrayList) Data() []int {
	return a.data
}

func (a *ArrayList) IsEmpty() bool {
	return a.Length() == 0
}

func (a *ArrayList) AddLast(value int) error {
	if err := a.Add(value, a.Length()+1); err != nil {
		return arrayIsFull
	}

	return nil
}

func (a *ArrayList) AddFirst(value int) error {
	if err := a.Add(value, 1); err != nil {
		return arrayIsFull
	}

	return nil
}

func (a *ArrayList) Add(value, index int) error {

	if index < 0 || index-2 > a.capacity {
		return invalidIndex
	}

	if a.size == a.Capacity() {
		if !a.allowResize {
			return arrayIsFull
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

func (a *ArrayList) Resize() {
	tmpData := make([]int, a.size, a.capacity/2)
	copy(tmpData, a.data[:a.size])
	a.data = tmpData
	a.capacity /= 2
}

func (a *ArrayList) Remove(index int) (int, error) {
	if index < 0 || index > a.size {
		return 0, invalidIndex
	}

	if a.size == 0 {
		return 0, arrayIsEmpty
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
			a.Resize()
		}
	}

	return ret, nil

}

func (a *ArrayList) RemoveFirst() (int, error) {
	v, err := a.Remove(1)
	if err != nil {
		return 0, arrayIsEmpty
	}

	return v, nil
}

func (a *ArrayList) RemoveLast() (int, error) {
	v, err := a.Remove(a.size)
	if err != nil {
		return 0, arrayIsEmpty
	}

	return v, nil
}

func (a *ArrayList) RemoveElement(value int) error {
	index := a.Find(value)
	if index != -1 {
		_, err := a.Remove(index + 1)
		return err
	}

	return nil
}

func (a *ArrayList) Get(index int) (int, error) {
	if index < 0 || index > a.size {
		return -1, invalidIndex
	}

	return a.data[index-1], nil
}

func (a *ArrayList) Set(value, index int) error {
	if index < 0 || index > a.size {
		return invalidIndex
	}

	a.data[index-1] = value

	return nil
}

func (a *ArrayList) Contains(value int) bool {
	if a.size == 0 {
		return false
	}

	for _, v := range a.data {
		if v == value {
			return true
		}
	}

	return false
}

func (a *ArrayList) Find(value int) int {
	if a.size == 0 {
		return -1
	}

	for i := 0; i < a.size; i++ {
		if a.data[i] == value {
			return i + 1
		}
	}

	return -1
}
