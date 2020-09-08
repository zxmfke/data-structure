package stack

import "testing"

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

func TestArrayStack_Push(t *testing.T) {
	type Cases struct {
		name      string
		value     int
		wantArray []int
		wantErr   error
	}

	arrStack := NewArrayStack(5)

	cases := []Cases{
		Cases{
			name:      "push first",
			value:     10,
			wantArray: []int{10},
			wantErr:   nil,
		},
		Cases{
			name:      "push second",
			value:     20,
			wantArray: []int{10, 20},
			wantErr:   nil,
		},
		Cases{
			name:      "push third",
			value:     30,
			wantArray: []int{10, 20, 30},
			wantErr:   nil,
		},
		Cases{
			name:      "push forth",
			value:     40,
			wantArray: []int{10, 20, 30, 40},
			wantErr:   nil,
		},
		Cases{
			name:      "push fifth",
			value:     50,
			wantArray: []int{10, 20, 30, 40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "stack is full",
			value:     60,
			wantArray: []int{10, 20, 30, 40, 50},
			wantErr:   stackIsFull,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			err := arrStack.Enqueue(v.value)
			if err != v.wantErr {
				t.Errorf("test stack push fail , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if isEqual := testEq(v.wantArray, arrStack.data); !isEqual {
						t.Errorf("test stack push index result is not equal \n expected : %v \n result   : %v", v.wantArray, arrStack.data)
					} else {
						t.Logf("%s", arrStack.String())
					}
				}
			}
		})
	}

	t.Log("finished")
}

func TestArrayStack_Pop(t *testing.T) {
	type Cases struct {
		name         string
		wantPopValue int
		wantErr      error
	}

	arrStack := NewArrayStack(5).WithData([]int{10, 20, 30, 40, 50})

	cases := []Cases{
		Cases{
			name:         "pop first",
			wantPopValue: 50,
			wantErr:      nil,
		},
		Cases{
			name:         "pop second",
			wantPopValue: 40,
			wantErr:      nil,
		},
		Cases{
			name:         "pop third",
			wantPopValue: 30,
			wantErr:      nil,
		},
		Cases{
			name:         "pop forth",
			wantPopValue: 20,
			wantErr:      nil,
		},
		Cases{
			name:         "pop fifth",
			wantPopValue: 10,
			wantErr:      nil,
		},
		Cases{
			name:         "stack is empty",
			wantPopValue: -1,
			wantErr:      stackIsEmpty,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			value, err := arrStack.Dequeue()
			if err != v.wantErr {
				t.Errorf("test stack pop fail , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if value != v.wantPopValue {
						t.Errorf("test stack pop result is not wrong \n expected : %v \n result   : %v", v.wantPopValue, arrStack.data)
					}
				}
			}
		})
	}

	t.Log("finished")
}
