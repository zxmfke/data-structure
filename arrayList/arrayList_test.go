package arrayList

import "testing"

//all test without resize option

func testEq(a, b []int) bool {
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

func TestNewArrayList(t *testing.T) {
	arr := NewArrayList(10)

	t.Logf("new array list with capacity, %s", arr.String())
}

func TestNewDefaultArrayList(t *testing.T) {

	arr := NewDefaultArrayList()

	t.Logf("new default array list %s", arr.String())
}

func TestArrayList_AddLast(t *testing.T) {
	type Cases struct {
		name      string
		value     int
		wantArray []int
		wantErr   error
	}

	arr := NewArrayList(5)

	cases := []Cases{
		Cases{
			name:      "insert first index",
			value:     10,
			wantArray: []int{10},
			wantErr:   nil,
		},
		Cases{
			name:      "insert second index",
			value:     20,
			wantArray: []int{10, 20},
			wantErr:   nil,
		},
		Cases{
			name:      "insert third index",
			value:     30,
			wantArray: []int{10, 20, 30},
			wantErr:   nil,
		},
		Cases{
			name:      "insert forth index",
			value:     40,
			wantArray: []int{10, 20, 30, 40},
			wantErr:   nil,
		},
		Cases{
			name:      "insert fifth index",
			value:     50,
			wantArray: []int{10, 20, 30, 40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "array is full",
			value:     60,
			wantArray: []int{10, 20, 30, 40, 50},
			wantErr:   arrayIsFull,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			err := arr.AddLast(v.value)
			if err != v.wantErr {
				t.Errorf("test add last fail , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if isEqual := testEq(v.wantArray, arr.data); !isEqual {
						t.Errorf("test add last result is not equal \n expected : %v \n result : %v", v.wantArray, arr.data)
					} else {
						t.Logf("%s", arr.String())
					}
				}
			}
		})
	}

	t.Log("finished")
}

func TestArrayList_AddFirst(t *testing.T) {
	type Cases struct {
		name      string
		value     int
		wantArray []int
		wantErr   error
	}

	arr := NewArrayList(5)

	cases := []Cases{
		Cases{
			name:      "insert first index",
			value:     10,
			wantArray: []int{10},
			wantErr:   nil,
		},
		Cases{
			name:      "insert second index",
			value:     20,
			wantArray: []int{20, 10},
			wantErr:   nil,
		},
		Cases{
			name:      "insert third index",
			value:     30,
			wantArray: []int{30, 20, 10},
			wantErr:   nil,
		},
		Cases{
			name:      "insert forth index",
			value:     40,
			wantArray: []int{40, 30, 20, 10},
			wantErr:   nil,
		},
		Cases{
			name:      "insert fifth index",
			value:     50,
			wantArray: []int{50, 40, 30, 20, 10},
			wantErr:   nil,
		},
		Cases{
			name:      "array is full",
			value:     60,
			wantArray: []int{50, 40, 30, 20, 10},
			wantErr:   arrayIsFull,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			err := arr.AddFirst(v.value)
			if err != v.wantErr {
				t.Errorf("test add first fail , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if isEqual := testEq(v.wantArray, arr.data); !isEqual {
						t.Errorf("test add first result is not equal \n expected : %v \n result : %v", v.wantArray, arr.data)
					} else {
						t.Logf("%s", arr.String())
					}
				}
			}
		})
	}

	t.Log("finished")
}

func TestArrayList_Add(t *testing.T) {
	type Cases struct {
		name      string
		value     int
		index     int
		wantArray []int
		wantErr   error
	}

	arr := NewArrayList(5)

	cases := []Cases{
		Cases{
			name:      "insert first index",
			value:     10,
			index:     1,
			wantArray: []int{10},
			wantErr:   nil,
		},
		Cases{
			name:      "insert second index",
			value:     20,
			index:     2,
			wantArray: []int{10, 20},
			wantErr:   nil,
		},
		Cases{
			name:      "insert third index",
			value:     30,
			index:     2,
			wantArray: []int{10, 30, 20},
			wantErr:   nil,
		},
		Cases{
			name:      "insert forth index",
			value:     40,
			index:     3,
			wantArray: []int{10, 30, 40, 20},
			wantErr:   nil,
		},
		Cases{
			name:      "insert fifth index",
			value:     50,
			index:     1,
			wantArray: []int{50, 10, 30, 40, 20},
			wantErr:   nil,
		},
		Cases{
			name:      "array is full",
			value:     60,
			index:     5,
			wantArray: []int{50, 10, 30, 40, 20},
			wantErr:   arrayIsFull,
		},
		Cases{
			name:      "invalid index insert",
			value:     60,
			index:     6,
			wantArray: []int{50, 10, 30, 40, 20},
			wantErr:   invalidIndex,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			err := arr.Add(v.value, v.index)
			if err != v.wantErr {
				t.Errorf("test add with index fail , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if isEqual := testEq(v.wantArray, arr.data); !isEqual {
						t.Errorf("test add with index result is not equal \n expected : %v \n result   : %v", v.wantArray, arr.data)
					} else {
						t.Logf("%s", arr.String())
					}
				}
			}
		})
	}

	t.Log("finished")
}

func TestArrayList_RemoveLast(t *testing.T) {
	type Cases struct {
		name      string
		wantArray []int
		wantErr   error
	}

	arr := NewDefaultArrayList().WithData([]int{10, 20, 30, 40, 50})

	cases := []Cases{
		Cases{
			name:      "#1 remove last index",
			wantArray: []int{10, 20, 30, 40},
			wantErr:   nil,
		},
		Cases{
			name:      "#2 remove last index",
			wantArray: []int{10, 20, 30},
			wantErr:   nil,
		},
		Cases{
			name:      "#3 remove last index",
			wantArray: []int{10, 20},
			wantErr:   nil,
		},
		Cases{
			name:      "#4 remove last index",
			wantArray: []int{10},
			wantErr:   nil,
		},
		Cases{
			name:      "#5 remove last index",
			wantArray: []int{},
			wantErr:   nil,
		},
		Cases{
			name:      "#6 array is empty",
			wantArray: []int{},
			wantErr:   arrayIsEmpty,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			_, err := arr.RemoveLast()
			if err != v.wantErr {
				t.Errorf("test remove last has error , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if isEqual := testEq(v.wantArray, arr.data); !isEqual {
						t.Errorf("test remove last result is not equal \n expected : %v \n result : %v", v.wantArray, arr.data)
					} else {
						t.Logf("%s", arr.String())
					}
				}
			}
		})
	}

	t.Log("finished")
}

func TestArrayList_RemoveFirst(t *testing.T) {
	type Cases struct {
		name      string
		wantArray []int
		wantErr   error
	}

	arr := NewDefaultArrayList().WithData([]int{10, 20, 30, 40, 50})

	cases := []Cases{
		Cases{
			name:      "#1 remove first index",
			wantArray: []int{20, 30, 40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "#2 remove first index",
			wantArray: []int{30, 40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "#3 remove first index",
			wantArray: []int{40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "#4 remove first index",
			wantArray: []int{50},
			wantErr:   nil,
		},
		Cases{
			name:      "#5 remove first index",
			wantArray: []int{},
			wantErr:   nil,
		},
		Cases{
			name:      "#6 array is empty",
			wantArray: []int{},
			wantErr:   arrayIsEmpty,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			_, err := arr.RemoveFirst()
			if err != v.wantErr {
				t.Errorf("test remove first has error , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if isEqual := testEq(v.wantArray, arr.data); !isEqual {
						t.Errorf("test remove first result is not equal \n expected : %v \n result : %v", v.wantArray, arr.data)
					} else {
						t.Logf("%s", arr.String())
					}
				}
			}
		})
	}

	t.Log("finished")
}

func TestArrayList_RemoveElement(t *testing.T) {
	type Cases struct {
		name      string
		wantArray []int
		value     int
		wantErr   error
	}

	arr := NewDefaultArrayList().WithData([]int{10, 20, 30, 40, 50})

	cases := []Cases{
		Cases{
			name:      "remove 10",
			value:     10,
			wantArray: []int{20, 30, 40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "remove 30",
			value:     30,
			wantArray: []int{20, 40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "remove 20",
			value:     20,
			wantArray: []int{40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "remove 40",
			value:     40,
			wantArray: []int{50},
			wantErr:   nil,
		},
		Cases{
			name:      "remove 50",
			value:     50,
			wantArray: []int{},
			wantErr:   nil,
		},
		Cases{
			name:      "remove 70",
			value:     70,
			wantArray: []int{},
			wantErr:   nil,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			err := arr.RemoveElement(v.value)
			if err != v.wantErr {
				t.Errorf("test remove element has error , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if isEqual := testEq(v.wantArray, arr.data); !isEqual {
						t.Errorf("test remove element result is not equal \n expected : %v \n result : %v", v.wantArray, arr.data)
					} else {
						t.Logf("%s", arr.String())
					}
				}
			}
		})
	}

	t.Log("finished")
}

func TestArrayList_Remove(t *testing.T) {
	type Cases struct {
		name      string
		index     int
		wantArray []int
		wantErr   error
	}

	arr := NewDefaultArrayList().WithData([]int{10, 20, 30, 40, 50})

	cases := []Cases{
		Cases{
			name:      "#1 remove second index",
			index:     2,
			wantArray: []int{10, 30, 40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "#2 remove second index",
			index:     2,
			wantArray: []int{10, 40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "#3 remove first index",
			index:     1,
			wantArray: []int{40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "#4 remove second index",
			index:     2,
			wantArray: []int{40},
			wantErr:   nil,
		},
		Cases{
			name:      "#5 remove invalid index",
			index:     2,
			wantArray: []int{40},
			wantErr:   invalidIndex,
		},
		Cases{
			name:      "#6 remove first index",
			index:     1,
			wantArray: []int{},
			wantErr:   nil,
		},
		Cases{
			name:      "#7 array is empty",
			index:     0,
			wantArray: []int{},
			wantErr:   arrayIsEmpty,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			_, err := arr.Remove(v.index)
			if err != v.wantErr {
				t.Errorf("test array remove fail , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if isEqual := testEq(v.wantArray, arr.data); !isEqual {
						t.Errorf("test array remove result is not equal \n expected : %v \n result   : %v", v.wantArray, arr.data)
					} else {
						t.Logf("%s", arr.String())
					}
				}
			}
		})
	}

	t.Logf("finised")
}

func TestArrayList_Find(t *testing.T) {
	type Cases struct {
		name      string
		value     int
		wantIndex int
	}

	arr := NewDefaultArrayList().WithData([]int{10, 20, 30, 40, 50})

	cases := []Cases{
		Cases{
			name:      "find 20",
			value:     20,
			wantIndex: 2,
		},
		Cases{
			name:      "find 40",
			value:     40,
			wantIndex: 4,
		},
		Cases{
			name:      "find 90",
			value:     90,
			wantIndex: -1,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			index := arr.Find(v.value)
			if index != v.wantIndex {
				t.Errorf("test array find result wrong , expected : %d, result : %d", v.wantIndex, index)
			}
		})
	}

	t.Logf("finised")
}

func TestArrayList_Contains(t *testing.T) {
	type Cases struct {
		name        string
		value       int
		wantContain bool
	}

	arr := NewDefaultArrayList().WithData([]int{10, 20, 30, 40, 50})

	cases := []Cases{
		Cases{
			name:        "find 20",
			value:       20,
			wantContain: true,
		},
		Cases{
			name:        "find 40",
			value:       40,
			wantContain: true,
		},
		Cases{
			name:        "find 90",
			value:       90,
			wantContain: false,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			has := arr.Contains(v.value)
			if has != v.wantContain {
				t.Errorf("test array find result wrong , expected : %v, result : %v", v.wantContain, has)
			}
		})
	}

	t.Logf("finised")
}

func TestArrayList_Capacity(t *testing.T) {
	type Cases struct {
		name         string
		array        *ArrayList
		wantCapacity int
	}

	cases := []Cases{
		Cases{
			name:         "capacity 5",
			array:        NewArrayList(5),
			wantCapacity: 5,
		},
		Cases{
			name:         "capacity 2",
			array:        NewArrayList(2),
			wantCapacity: 2,
		},
		Cases{
			name:         "capacity 5",
			array:        NewArrayList(5),
			wantCapacity: 5,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			_ = v.array.AddLast(1)

			if v.wantCapacity != v.array.Capacity() {
				t.Errorf("test array capacity result wrong , expected : %d, result : %d", v.wantCapacity, v.array.Capacity())
			}
		})
	}

	t.Logf("finished")
}

func TestArrayList_IsEmpty(t *testing.T) {
	type Cases struct {
		name         string
		array        *ArrayList
		isNeedInsert bool
		wantEmpty    bool
	}

	cases := []Cases{
		Cases{
			name:         "array empty",
			array:        NewArrayList(5),
			isNeedInsert: false,
			wantEmpty:    true,
		},
		Cases{
			name:         "array not empty",
			array:        NewArrayList(2),
			isNeedInsert: true,
			wantEmpty:    false,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			if v.isNeedInsert {
				_ = v.array.AddLast(1)
			}

			if v.wantEmpty != v.array.IsEmpty() {
				t.Errorf("test array is empty result wrong , expected : %v, result : %v", v.wantEmpty, v.array.IsEmpty())
			}
		})
	}

	t.Logf("finished")
}

func TestArrayList_Get(t *testing.T) {
	type Cases struct {
		name      string
		index     int
		wantValue int
		wantErr   error
	}

	arr := NewDefaultArrayList().WithData([]int{10, 20, 30, 40, 50})

	cases := []Cases{
		Cases{
			name:      "get the second",
			index:     2,
			wantValue: 20,
			wantErr:   nil,
		},
		Cases{
			name:      "get the fifth",
			index:     5,
			wantValue: 50,
			wantErr:   nil,
		},
		Cases{
			name:      "get invalid index",
			index:     9,
			wantValue: 50,
			wantErr:   invalidIndex,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			value, err := arr.Get(v.index)
			if err != v.wantErr {
				t.Errorf("test array get error , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if v.wantValue != value {
						t.Errorf("test array get result wrong , expected : %d, result : %d", v.wantValue, value)
					}
				}
			}
		})
	}

	t.Logf("finised")
}

func TestArrayList_Set(t *testing.T) {
	type Cases struct {
		name      string
		index     int
		value     int
		wantArray []int
		wantErr   error
	}

	arr := NewDefaultArrayList().WithData([]int{10, 20, 30, 40, 50})

	cases := []Cases{
		Cases{
			name:      "set the second to 50",
			index:     2,
			value:     50,
			wantArray: []int{10, 50, 30, 40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "set the firth to 40",
			index:     1,
			value:     40,
			wantArray: []int{40, 50, 30, 40, 50},
			wantErr:   nil,
		},
		Cases{
			name:      "set invalid index",
			index:     9,
			value:     40,
			wantArray: []int{40, 50, 30, 40, 50},
			wantErr:   invalidIndex,
		},
	}

	for _, v := range cases {

		t.Run(v.name, func(t *testing.T) {
			err := arr.Set(v.value, v.index)
			if err != v.wantErr {
				t.Errorf("test array set error , expected : %v, result : %v", v.wantErr, err)
			} else {
				if err == nil {
					if isEqual := testEq(v.wantArray, arr.data); !isEqual {
						t.Errorf("test array set result is not equal \n expected : %v \n result   : %v", v.wantArray, arr.data)
					} else {
						t.Logf("%s", arr.String())
					}
				}
			}
		})
	}

	t.Logf("finised")
}
