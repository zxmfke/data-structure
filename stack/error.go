package stack

import "github.com/pkg/errors"

var (
	invalidIndex = errors.New("invalid index, require index >=0 && index <= arrayStack length, array stack add fail")
	stackIsFull  = errors.New("stack is full, add fail")
	stackIsEmpty = errors.New("stack is empty, add fail")
)