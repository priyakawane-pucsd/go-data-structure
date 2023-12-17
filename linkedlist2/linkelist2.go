package linkedlist2

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type SinglyLinkeList struct {
	start *Node
	end   *Node
}

func newNode(value int) *Node {
	return &Node{value: value}
}

func NewSinglyLinkeList() *SinglyLinkeList {
	return &SinglyLinkeList{}
}

// method belongs to that object only, we cannot call that method
func (sl *SinglyLinkeList) SortedAdd(value int) {
	node := newNode(value)
	// if linklist is empty
	if sl.start == nil {
		sl.start = node
		sl.end = node
		return
	}

	//if value is less than first node
	if value < sl.start.value {
		temp := sl.start
		sl.start = node
		sl.start.next = temp
		return
	}

	//1->3->4->7->8->nil
	//value is lies in between of linkelist
	temp := sl.start
	for temp.next != nil {
		if value < temp.next.value {
			node.next = temp.next
			temp.next = node
			return
		}
		temp = temp.next
	}

	//if value needs to add at end
	sl.end.next = node
	sl.end = node
	return
}

// 1->2->3->8->10->nil
func (sl *SinglyLinkeList) Remove(value int) {
	//if value present at start node
	if value == sl.start.value {
		sl.start = sl.start.next
		return
	}

	//if value present in between of the node
	temp := sl.start
	for temp.next != nil {
		if value == temp.next.value {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
}

func (sl *SinglyLinkeList) Display() {
	current := sl.start
	for current != nil {
		fmt.Printf("%d->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
