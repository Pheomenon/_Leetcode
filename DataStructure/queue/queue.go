package queue

type node struct {
	pre   *node
	next  *node
	value interface{}
}

type Queue struct {
	top    *node
	rear   *node
	length int
}

func New() *Queue {
	return &Queue{
		top:    nil,
		rear:   nil,
		length: 0,
	}
}

func (q *Queue) Length() int {
	return q.length
}

//返回true队列不为空
func (q *Queue) Any() bool {
	if q.length != 0 {
		return true
	}
	return false
}

//返回队列顶端元素
func (q *Queue) Peek() interface{} {
	if q.top == nil {
		return nil
	}
	return q.top.value
}

//入队操作
func (q *Queue) Push(v interface{}) {
	newNode := &node{
		pre:   q.rear,
		next:  nil,
		value: v,
	}
	if q.length == 0 {
		q.top = newNode
		q.rear = newNode
	} else {
		newNode.pre = q.rear
		q.rear.next = newNode
		q.rear = newNode
	}
	q.length++
}

//出队操作
func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}
	top := q.top
	if q.top.next == nil {
		q.top = nil
	} else {
		q.top = q.top.next
		q.top.pre.next = nil
		q.top.pre = nil
	}
	q.length--
	return top.value
}
