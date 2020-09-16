package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewMaxHeap(t *testing.T) {
	arrayHeap := NewDefaultMaxHeap()

	rand.Seed(time.Now().UnixNano())

	tmpMap := make(map[int]bool, 0)

	//tmpArray := []int{62,41,30,28,16,22,13,19,17,15,52}

	for i := 0; i < 50; i++ {
		v := rand.Intn(50)
		if _, ok := tmpMap[v]; !ok {
			arrayHeap.Add(v)
			tmpMap[v] = true
		} else {
			i--
		}
	}

	ret := ""
	for i := 0; i < 50; i++ {
		v, err := arrayHeap.ExtractMax()
		if err != nil {
			t.Fatalf("%s", err.Error())
		}
		ret = fmt.Sprintf("%s,%d", ret, v)
	}

	t.Logf("result : %s", ret)
}
