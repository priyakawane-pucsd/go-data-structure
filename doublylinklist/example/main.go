package main

import (
	"fmt"
	"go-data-structure/doublylinklist"
)

func main() {
	doublylist := doublylinklist.NewDoublyLinkList()
	// doublylist.Add(3)
	// doublylist.Add(4)
	// doublylist.Add(8)
	doublylist.SortedAdd(8)
	doublylist.SortedAdd(2)
	doublylist.SortedAdd(5)
	fmt.Println("1. Disply Doubly linked list:-")
	doublylist.Display()
	fmt.Println("2. Disply Doubly linked list in reverse way:-")
	doublylist.DisplayReverse()
	doublylist.Remove(2)
	fmt.Println("3. Disply Doubly linked list in reverse way after removal****:-")
	doublylist.DisplayReverse()
	fmt.Println("Disply Doubly linked list after removal of node:-")
	doublylist.Display()
}
