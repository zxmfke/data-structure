package queue

import (
	"math/rand"
	"testing"
	"time"
)

func TestLinkedQueue_Enqueue(t *testing.T) {
	type Cases struct {
		name      string
		value     int
		wantArray []int
		wantErr   error
	}

	linkedListQueue := NewLinkedListQueue()

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
			_ = linkedListQueue.Enqueue(v.value)
			t.Logf("%s", linkedListQueue.String())
		})
	}

	t.Log("finished")
}

func TestLinkedQueue_Dequeue(t *testing.T) {
	type Cases struct {
		name         string
		wantPopValue int
		wantErr      error
	}

	linkedListQueue := NewLinkedListQueue()

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
	}

	for _, v := range cases {
		_ = linkedListQueue.Enqueue(v.wantPopValue)
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			value, err := linkedListQueue.Dequeue()
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

func BenchmarkLinkedListQueue_Enqueue(b *testing.B) {
	var times = 10000
	linkedListQueue := NewLinkedListQueue()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		v := rand.Intn(100)
		_ = linkedListQueue.Enqueue(v)
	}
}

func BenchmarkLinkedListQueue_Dequeue(b *testing.B) {
	var times = 10000
	linkedListQueue := NewLinkedListQueue()
	for i := 0; i < times; i++ {
		v := rand.Intn(100)
		_ = linkedListQueue.Enqueue(v)
	}

	b.ResetTimer()

	for i := 0; i < times; i++ {
		_, _ = linkedListQueue.Dequeue()
	}
}
