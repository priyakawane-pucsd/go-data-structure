package main

import (
	"fmt"
	"go-data-structure/linkedlist2"
)

func main() {
	singlyLinkList := linkedlist2.NewSinglyLinkeList()
	singlyLinkList.SortedAdd(2)
	singlyLinkList.SortedAdd(3)
	singlyLinkList.SortedAdd(8)
	singlyLinkList.SortedAdd(1)
	singlyLinkList.SortedAdd(10)
	fmt.Println("Linkedlist after insertion:-")
	singlyLinkList.Display()
	singlyLinkList.Remove(10)
	fmt.Println("Linkedlist after removal of node 3:-")
	singlyLinkList.Display()
}
