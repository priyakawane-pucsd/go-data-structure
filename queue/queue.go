package queue

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	start *Node
	end   *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func NewQueueNode(value int) *Node {
	return &Node{value: value}
}

func (q *Queue) EnQueue(value int) {
	node := NewQueueNode(value)
	if q.start == nil {
		q.start = node
		q.end = node
		return
	}
	q.end.next = node
	q.end = node
	return
}

// 4->5->7
func (q *Queue) DeQueue() {
	q.start = q.start.next
	return
}

func (q *Queue) Display() {
	current := q.start
	for current != nil {
		fmt.Printf("%d ->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
