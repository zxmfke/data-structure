package arrayList

import "github.com/pkg/errors"

var (
	invalidIndex = errors.New("invalid index, require index >=0 && index <= arrayList length, array list add fail")
	arrayIsFull = errors.New("array is full, add fail")
	arrayIsEmpty = errors.New("array is empty, add fail")
)