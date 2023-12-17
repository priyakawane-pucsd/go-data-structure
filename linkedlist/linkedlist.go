package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

type SinglyLinkeList struct {
	start *Node
	end   *Node
}

func NewSinglyLinkeList() *SinglyLinkeList {
	return &SinglyLinkeList{}
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

func (sl *SinglyLinkeList) Add(value int) {
	node := NewNode(value)
	//check if in current list there is no node, then set start with new node & end also points to the new node
	//else the current end points to new node & new node become the end
	if sl.start == nil {
		sl.start = node
		sl.end = node
		return
	}
	sl.end.next = node
	sl.end = node
	return
}

func (sl *SinglyLinkeList) Display() {
	current := sl.start
	for current != nil {
		fmt.Printf("%d ->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func (sl *SinglyLinkeList) Remove(val any) {
	if sl.start == nil {
		return
	}
	// remove first node if its matches
	if sl.start.value == val {
		start := sl.start
		sl.start = start.next
		if start == sl.end {
			sl.end = start.next
		}
		return
	}

	//1->2->3->4->5->nil
	temp := sl.start
	for temp.next != nil {
		if temp.next.value == val {
			temp.next = temp.next.next
			return
		} else {
			temp = temp.next
		}
	}
}

// 4->6->10
func (sl *SinglyLinkeList) SortedAdd(value int) {
	newNode := NewNode(value)
	// if linklist is empty
	if sl.start == nil {
		sl.start = newNode
		sl.end = newNode
		return
	}
	// if node should add at start
	if value < sl.start.value {
		newNode.next = sl.start
		sl.start = newNode
		return
	}

	// if node should store in between
	temp := sl.start
	for temp.next != nil {
		if temp.next.value > value {
			//store here
			newNode.next = temp.next
			temp.next = newNode
			return
		}
		temp = temp.next
	}

	// if needs to add at end
	temp.next = newNode
	sl.end = newNode
}
