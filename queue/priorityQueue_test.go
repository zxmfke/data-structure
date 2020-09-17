package queue

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPriorityQueue_Enqueue(t *testing.T) {
	pQ := NewPriorityQueue()
	tmpMap := make(map[int]bool, 0)

	for i := 0; i < 50; i++ {
		v := rand.Intn(50)
		if _, ok := tmpMap[v]; !ok {
			_ = pQ.Enqueue(v)
			tmpMap[v] = true
		} else {
			i--
		}
	}
}

func TestPriorityQueue_Dequeue(t *testing.T) {
	pQ := NewPriorityQueue()
	tmpMap := make(map[int]bool, 0)

	for i := 0; i < 50; i++ {
		v := rand.Intn(50)
		if _, ok := tmpMap[v]; !ok {
			_ = pQ.Enqueue(v)
			tmpMap[v] = true
		} else {
			i--
		}
	}

	ret := ""
	for i := 0; i < 50; i++ {
		v, err := pQ.Dequeue()
		if err != nil {
			t.Fatalf("%s", err.Error())
		}
		ret = fmt.Sprintf("%s,%d", ret, v.(int))
	}

	t.Logf("result : %s", ret)
}
