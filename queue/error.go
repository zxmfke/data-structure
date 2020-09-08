package queue

import "github.com/pkg/errors"

var (
	invalidIndex = errors.New("invalid index, require index >=0 && index <= arrayQueue length, array queue enqueue fail")
	queueIsFull  = errors.New("queue is full")
	queueIsEmpty = errors.New("queue is empty")
)