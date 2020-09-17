package heap

import "github.com/navi-tt/data-structure/arrayList"

/**
 * @Author: Zheng xiaomin
 * @Date: 2020/9/17 22:00 晚上
 */


//Add       O(n)   因为添加元素有可能导致扩容
//RemoveMax O(logn)

//insert with array
//by add     O(nlogn)
//by heapify O(n)

//可以添加一个比较器，来定义比较的方式，不用单单比较int的大小

//leecode 347

//扩展
// d-ary Heap
// 索引堆
// 二项堆
// 斐波那契堆

type MaxHeap struct {
	arr *arrayList.ArrayList
}

func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{
		arr: arrayList.NewArrayList(capacity).WithResize(),
	}
}

func NewDefaultMaxHeap() *MaxHeap {
	return &MaxHeap{
		arr: arrayList.NewDefaultArrayList().WithResize(),
	}
}

// O(logn)
func NewMaxHeapByHeapify(arr []int) *MaxHeap {

	maxHeap := NewMaxHeap(len(arr)).WithData(arr)

	for i := maxHeap.parent(len(arr) - 1); i >= 0; i-- {
		maxHeap.siftDown(i)
	}

	return maxHeap
}

func (m *MaxHeap) WithData(arr []int) *MaxHeap {
	m.arr.WithData(arr)
	return m
}

func (m *MaxHeap) GetSize() int {
	return m.arr.Length()
}

func (m *MaxHeap) Capacity() int {
	return m.arr.Capacity()
}

func (m *MaxHeap) IsEmpty() bool {
	return m.arr.IsEmpty()
}

func (m *MaxHeap) FindMax() (int, error) {
	if m.GetSize() == 0 {
		return -1, heapIsEmpty
	}
	return m.get(0), nil
}

func (m *MaxHeap) get(index int) int {
	v, _ := m.arr.Get(index + 1)
	return v
}

func (m *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}

func (m *MaxHeap) leftChild(index int) int {
	return index*2 + 1
}

func (m *MaxHeap) rightChild(index int) int {
	return index*2 + 2
}

func (m *MaxHeap) Add(value int) {
	_ = m.arr.AddLast(value)
	m.siftUp(m.GetSize() - 1)
}

func (m *MaxHeap) siftUp(index int) {
	for ; ; index = m.parent(index) {
		if index == 0 {
			return
		}
		parent := m.get(m.parent(index))
		cur := m.get(index)
		if parent < cur {
			_ = m.arr.Swap(m.parent(index), index)
		}
	}
}

func (m *MaxHeap) ExtractMax() (int, error) {
	max, err := m.getMax()
	if err != nil {
		return -1, err
	}

	_ = m.arr.Swap(0, m.GetSize()-1)
	_, err = m.arr.RemoveLast()
	if err != nil {
		return -1, err
	}

	m.siftDown(0)
	return max, nil
}

func (m *MaxHeap) getMax() (int, error) {
	if m.GetSize() == 0 {
		return -1, heapIsEmpty
	}

	return m.arr.Get(1)
}

func (m *MaxHeap) siftDown(index int) {
	for ; m.leftChild(index) < m.GetSize(); {
		j := m.leftChild(index)

		if j+1 < m.GetSize() && m.get(j) < m.get(m.rightChild(index)) {
			j = m.rightChild(index)
		}

		if m.get(index) > m.get(j) {
			return
		}

		_ = m.arr.Swap(index, j)
		index = j
	}
}

// O(logn)
func (m *MaxHeap) Replace(value int) (int, error) {
	maxValue, err := m.FindMax()
	if err != nil {
		return -1, err
	}

	if err := m.arr.Set(value, 0); err != nil {
		return -1, err
	}

	m.siftDown(0)
	return maxValue, nil
}
