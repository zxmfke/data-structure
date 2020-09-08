package queue

import (
	"math/rand"
	"testing"
	"time"
)

func TestLoopQueue_Enqueue(t *testing.T) {
	type Cases struct {
		name      string
		value     int
		wantArray []int
		wantErr   error
	}

	loopQueue := NewLoopQueue(5)

	cases := []Cases{
		Cases{
			name:      "enqueue first",
			value:     10,
			wantArray: []int{10},
			wantErr:   nil,
		},
		Cases{
			name:      "enqueue second",
			value:     20,
			wantArray: []int{10, 20},
			wantErr:   nil,
		},
		Cases{
			name:      "enqueue third",
			value:     30,
			wantArray: []int{10, 20, 30},
			wantErr:   nil,
		},
		Cases{
			name:      "enqueue forth",
			value:     40,
			wantArray: []int{10, 20, 30, 40},
			wantErr:   nil,
		},
		Cases{
			name:      "enqueue fifth",
			value:     50,
			wantArray: []int{10, 20, 30, 40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "queue is full",
			value:     60,
			wantArray: []int{10, 20, 30, 40, 50},
			wantErr:   queueIsFull,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			err := loopQueue.Enqueue(v.value)
			if err != v.wantErr {
				t.Errorf("test queue enqueue fail , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if isEqual := testEq(v.wantArray, loopQueue.data); !isEqual {
						t.Errorf("test queue enqueue result is not equal \n expected : %v \n result   : %v", v.wantArray, loopQueue.data)
					} else {
						t.Logf("%s", loopQueue.String())
					}
				}
			}
		})
	}

	t.Log("finished")
}

func TestLoopQueue_Dequeue(t *testing.T) {
	type Cases struct {
		name         string
		wantPopValue int
		wantErr      error
	}

	loopQueue := NewArrayQueue(5).WithData([]int{10, 20, 30, 40, 50})

	cases := []Cases{
		Cases{
			name:         "dequeue first",
			wantPopValue: 10,
			wantErr:      nil,
		},
		Cases{
			name:         "dequeue second",
			wantPopValue: 20,
			wantErr:      nil,
		},
		Cases{
			name:         "dequeue third",
			wantPopValue: 30,
			wantErr:      nil,
		},
		Cases{
			name:         "dequeue forth",
			wantPopValue: 40,
			wantErr:      nil,
		},
		Cases{
			name:         "dequeue fifth",
			wantPopValue: 50,
			wantErr:      nil,
		},
		Cases{
			name:         "queue is empty",
			wantPopValue: -1,
			wantErr:      queueIsEmpty,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			value, err := loopQueue.Dequeue()
			if err != v.wantErr {
				t.Errorf("test queue dequeue fail , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if value != v.wantPopValue {
						t.Errorf("test queue dequeue result is not wrong \n expected : %v \n result   : %v", v.wantPopValue, value)
					}
				}
			}
		})
	}

	t.Log("finished")
}

func BenchmarkLoopQueue_Enqueue(b *testing.B) {
	var times = 10000
	loopQueue := NewLoopQueue(times)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		v := rand.Intn(100)
		_ = loopQueue.Enqueue(v)
	}
}

func BenchmarkLoopQueue_Dequeue(b *testing.B) {
	var times = 10000
	tmpLoopQueue := NewLoopQueue(times)
	for i := 0; i < times; i++ {
		v := rand.Intn(100)
		_ = tmpLoopQueue.Enqueue(v)
	}

	b.ResetTimer()

	for i := 0; i < times; i++ {
		_, _ = tmpLoopQueue.Dequeue()
	}
}
