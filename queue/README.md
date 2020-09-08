### Queue
---
Queue接口定义了

GetSize() int <br/>
IsEmpty() bool <br/>
Enqueue(interface{}) error <br/>
Dequeue() (interface{}, error) <br/>
GetFront() (interface{}, error) <br/>

目前有arrayQueue和loopQueue分别实现了Queue这个接口，这两个queue的区别在于，loopQueue的depueue的时间复杂度是O(1)，而arrayQueue的时间复杂度是O(n)
单元测试里面也有benchmark，差别挺大的。

BenchmarkLoopQueue_Dequeue   	1000000000	        0.000976 ns/op
BenchmarkArrayQueue_Dequeue   	1000000000	         0.0781 ns/op