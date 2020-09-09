package linkedList

import (
	"fmt"
	"testing"
)

func TestLinkedList_AddFirst(t *testing.T) {
	type Cases struct {
		name    string
		value   int
		wantErr error
	}

	List := NewLinkedList()

	cases := []Cases{
		Cases{
			name:    "insert first",
			value:   10,
			wantErr: nil,
		},
		Cases{
			name:    "insert second",
			value:   20,
			wantErr: nil,
		},
		Cases{
			name:    "insert third",
			value:   30,
			wantErr: nil,
		},
		Cases{
			name:    "insert forth",
			value:   40,
			wantErr: nil,
		},
		Cases{
			name:    "insert fifth",
			value:   50,
			wantErr: nil,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			List.AddFirst(v.value)
			fmt.Printf("%s", List)
		})
	}

	t.Log("finished")
}

func TestLinkedList_AddLast(t *testing.T) {
	type Cases struct {
		name    string
		value   int
		wantErr error
	}

	List := NewLinkedList()

	cases := []Cases{
		Cases{
			name:    "insert first",
			value:   10,
			wantErr: nil,
		},
		Cases{
			name:    "insert second",
			value:   20,
			wantErr: nil,
		},
		Cases{
			name:    "insert third",
			value:   30,
			wantErr: nil,
		},
		Cases{
			name:    "insert forth",
			value:   40,
			wantErr: nil,
		},
		Cases{
			name:    "insert fifth",
			value:   50,
			wantErr: nil,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			List.AddLast(v.value)
			fmt.Printf("%s", List)
		})
	}

	t.Log("finished")
}

func TestLinkedList_Add(t *testing.T) {
	type Cases struct {
		name    string
		value   int
		index   int
		wantErr error
	}

	List := NewLinkedList()

	cases := []Cases{
		Cases{
			name:    "insert first",
			value:   10,
			index:   1,
			wantErr: nil,
		},
		Cases{
			name:    "insert second",
			value:   20,
			index:   2,
			wantErr: nil,
		},
		Cases{
			name:    "insert third",
			value:   30,
			index:   3,
			wantErr: nil,
		},
		Cases{
			name:    "insert forth",
			value:   40,
			index:   2,
			wantErr: nil,
		},
		Cases{
			name:    "insert fifth",
			value:   50,
			index:   1,
			wantErr: nil,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			_ = List.Add(v.value, v.index)
			fmt.Printf("%s\n", List)
		})
	}

	t.Log("finished")
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	type Cases struct {
		name    string
		value   int
		wantErr error
	}

	List := NewLinkedList()

	cases := []Cases{
		Cases{
			name:    "insert first",
			value:   10,
			wantErr: nil,
		},
		Cases{
			name:    "insert second",
			value:   20,
			wantErr: nil,
		},
		Cases{
			name:    "insert third",
			value:   30,
			wantErr: nil,
		},
		Cases{
			name:    "insert forth",
			value:   40,
			wantErr: nil,
		},
		Cases{
			name:    "insert fifth",
			value:   50,
			wantErr: nil,
		},
	}

	for _, v := range cases {
		List.AddLast(v.value)
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			_, _ = List.RemoveFirst()
			fmt.Printf("%s", List)
		})
	}

	t.Log("finished")
}

func TestLinkedList_RemoveLast(t *testing.T) {
	type Cases struct {
		name    string
		value   int
		wantErr error
	}

	List := NewLinkedList()

	cases := []Cases{
		Cases{
			name:    "insert first",
			value:   10,
			wantErr: nil,
		},
		Cases{
			name:    "insert second",
			value:   20,
			wantErr: nil,
		},
		Cases{
			name:    "insert third",
			value:   30,
			wantErr: nil,
		},
		Cases{
			name:    "insert forth",
			value:   40,
			wantErr: nil,
		},
		Cases{
			name:    "insert fifth",
			value:   50,
			wantErr: nil,
		},
	}

	for _, v := range cases {
		List.AddLast(v.value)
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			_, _ = List.RemoveLast()
			fmt.Printf("%s", List)
		})
	}

	t.Log("finished")
}

func TestLinkedList_Remove(t *testing.T) {
	type Cases struct {
		name    string
		value   int
		index   int
		wantErr error
	}

	List := NewLinkedList()

	cases := []Cases{
		Cases{
			name:    "remote first",
			value:   10,
			index:   2,
			wantErr: nil,
		},
		Cases{
			name:    "remote second",
			value:   20,
			index:   1,
			wantErr: nil,
		},
		Cases{
			name:    "remote third",
			value:   30,
			index:   2,
			wantErr: nil,
		},
		Cases{
			name:    "remote forth",
			value:   40,
			index:   1,
			wantErr: nil,
		},
		Cases{
			name:    "remote fifth",
			value:   50,
			index:   1,
			wantErr: nil,
		},
	}

	for _, v := range cases {
		List.AddLast(v.value)
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			_, _ = List.Remove(v.index)
			fmt.Printf("%s", List)
		})
	}

	t.Log("finished")
}
