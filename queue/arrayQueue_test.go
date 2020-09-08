package queue

import (
	"math/rand"
	"testing"
	"time"
)

func testEq(a []int, b []interface{}) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestArrayQueue_Enqueue(t *testing.T) {
	type Cases struct {
		name      string
		value     int
		wantArray []int
		wantErr   error
	}

	arrQueue := NewArrayQueue(5)

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
			err := arrQueue.Enqueue(v.value)
			if err != v.wantErr {
				t.Errorf("test queue enqueue fail , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if isEqual := testEq(v.wantArray, arrQueue.data); !isEqual {
						t.Errorf("test queue enqueue result is not equal \n expected : %v \n result   : %v", v.wantArray, arrQueue.data)
					} else {
						t.Logf("%s", arrQueue.String())
					}
				}
			}
		})
	}

	t.Log("finished")
}

func TestArrayQueue_Dequeue(t *testing.T) {
	type Cases struct {
		name         string
		wantPopValue int
		wantErr      error
	}

	arrQueue := NewArrayQueue(5).WithData([]int{10, 20, 30, 40, 50})

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
			value, err := arrQueue.Dequeue()
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

func BenchmarkArrayQueue_Enqueue(b *testing.B) {
	var times = 10000
	arrayQueue := NewArrayQueue(times)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		v := rand.Intn(100)
		_ = arrayQueue.Enqueue(v)
	}
}

func BenchmarkArrayQueue_Dequeue(b *testing.B) {
	var times = 10000
	arrayQueue := NewArrayQueue(times)
	for i := 0; i < times; i++ {
		v := rand.Intn(100)
		_ = arrayQueue.Enqueue(v)
	}

	b.ResetTimer()

	for i := 0; i < times; i++ {
		_, _ = arrayQueue.Dequeue()
	}
}
