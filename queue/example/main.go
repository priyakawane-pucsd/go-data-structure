package main

import "go-data-structure/queue"

func main() {
	queue := queue.NewQueue()
	queue.EnQueue(4)
	queue.EnQueue(7)
	queue.EnQueue(8)
	queue.Display()
	queue.DeQueue()
	queue.DeQueue()
	queue.Display()
}
