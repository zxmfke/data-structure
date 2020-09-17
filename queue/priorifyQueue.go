package queue

import "github.com/navi-tt/data-structure/heap"

/**
 * @Author: Zheng xiaomin
 * @Date: 2020/9/17 22:00 晚上
 */

//Enqueue  O(n)
//Dequeue  O(1)
//GetFront O(1)

type PriorityQueue struct {
	maxHeap *heap.MaxHeap
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		maxHeap: heap.NewDefaultMaxHeap(),
	}
}

func NewPriorityQueueWithData(arr []int) *PriorityQueue {
	return &PriorityQueue{
		maxHeap: heap.NewMaxHeapByHeapify(arr),
	}
}

func (p *PriorityQueue) GetSize() int {
	return p.maxHeap.GetSize()
}

func (p *PriorityQueue) IsEmpty() bool {
	return p.maxHeap.IsEmpty()
}

func (p *PriorityQueue) GetFront() (interface{}, error) {
	return p.maxHeap.FindMax()
}

func (p *PriorityQueue) Enqueue(value interface{}) error {
	p.maxHeap.Add(value.(int))
	return nil
}

func (p *PriorityQueue) Dequeue() (interface{}, error) {
	return p.maxHeap.ExtractMax()
}
