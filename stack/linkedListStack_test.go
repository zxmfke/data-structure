package stack

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLinkedListStack_Push(t *testing.T) {
	type Cases struct {
		name    string
		value   int
		wantErr error
	}

	List := NewLinkedListStack()

	cases := []Cases{
		Cases{
			name:    "push first",
			value:   10,
			wantErr: nil,
		},
		Cases{
			name:    "push second",
			value:   20,
			wantErr: nil,
		},
		Cases{
			name:    "push third",
			value:   30,
			wantErr: nil,
		},
		Cases{
			name:    "push forth",
			value:   40,
			wantErr: nil,
		},
		Cases{
			name:    "push fifth",
			value:   50,
			wantErr: nil,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			_ = List.Push(v.value)
			fmt.Printf("%s", List)
		})
	}

	t.Log("finished")
}

func TestLinkedListStack_Pop(t *testing.T) {
	type Cases struct {
		name    string
		value   int
		wantErr error
	}

	List := NewLinkedListStack()

	cases := []Cases{
		Cases{
			name:    "pop first",
			value:   10,
			wantErr: nil,
		},
		Cases{
			name:    "pop second",
			value:   20,
			wantErr: nil,
		},
		Cases{
			name:    "pop third",
			value:   30,
			wantErr: nil,
		},
		Cases{
			name:    "pop forth",
			value:   40,
			wantErr: nil,
		},
		Cases{
			name:    "pop fifth",
			value:   50,
			wantErr: nil,
		},
	}

	for _, v := range cases {
		_ = List.Push(v.value)
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			value, _ := List.Pop()
			if value != v.value {
				t.Errorf("test stack pop result is not wrong \n expected : %v \n result   : %v", v.value, value)

			} else {
				fmt.Printf("%s", List)
			}
		})
	}

	t.Log("finished")
}

func BenchmarkLinkedListStack_Push(b *testing.B) {
	var times = 10000
	linkedListStack := NewLinkedListStack()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		v := rand.Intn(100)
		_ = linkedListStack.Push(v)
	}
}

func BenchmarkLinkedListStack_Pop(b *testing.B) {
	var times = 10000
	linkedListStack := NewLinkedListStack()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		v := rand.Intn(100)
		_ = linkedListStack.Push(v)
	}

	b.ResetTimer()

	for i := 0; i < times; i++ {
		_, _ = linkedListStack.Pop()
	}
}
