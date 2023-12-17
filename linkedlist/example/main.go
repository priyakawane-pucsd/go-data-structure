package main

import (
	"go-data-structure/linkedlist"
)

type Integer int

func (a Integer) Lessthan(i Integer) bool {
	return a < i
}

func main() {
	// singlyList := linkedlist.NewSinglyLinkeList()
	// singlyList.Add(8)
	// singlyList.Add(4)
	// singlyList.Add(5)
	// singlyList.Add(6)
	// singlyList.Display()
	// singlyList.Remove(6)
	// singlyList.Display()

	singlyList := linkedlist.NewSinglyLinkeList()
	singlyList.SortedAdd(11)
	singlyList.SortedAdd(4)
	singlyList.SortedAdd(5)
	//singlyList.Add(6)
	singlyList.Display()
	// singlyList.Remove(6)
	// singlyList.Display()

	//fmt.Println("SInglyLinkListtttttt", singlyList)
}
