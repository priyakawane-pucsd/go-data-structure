package doublylinklist

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type DoublyLinkList struct {
	start *Node
	end   *Node
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

func NewDoublyLinkList() *DoublyLinkList {
	return &DoublyLinkList{}
}

func (dl *DoublyLinkList) Add(value int) {
	node := NewNode(value)
	//fmt.Println("Node is>>>>", node)

	//if linkedlist is empty
	if dl.start == nil {
		dl.start = node
		dl.end = node
		return
	}
	// 3<-4
	dl.end.next = node
	nodePrev := dl.end //get end previous in one variable ie. nodePrev
	dl.end = node      //assign last node to end
	dl.end.prev = nodePrev
	nodePrev.next = dl.end
	return
}

func (dl *DoublyLinkList) Remove(value int) {
	//if value present at first position
	if value == dl.start.value {
		dl.start = dl.start.next
		dl.start.prev = nil
		return
	}

	//if value present at end
	if value == dl.end.value {
		dl.end = dl.end.prev
		dl.end.next = nil
		return
	}

	//if value is in between of list
	temp := dl.start
	for temp.next != nil {
		// 3->4->8
		if value == temp.next.value {
			temp.next = temp.next.next
			temp.next.prev = temp
			return
		}
		temp = temp.next
	}
}

func (dl *DoublyLinkList) Display() {
	current := dl.start
	for current != nil {
		fmt.Printf("%d->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

// 3->4->8
func (dl *DoublyLinkList) DisplayReverse() {
	current := dl.end
	for current != nil {
		fmt.Printf("%d->", current.value)
		current = current.prev
	}
	fmt.Println("nil")
}

func (dl *DoublyLinkList) SortedAdd(value int) {
	node := NewNode(value)

	//if doubly list is empty
	if dl.start == nil {
		dl.start = node
		dl.end = node
		return
	}

	//3->4->5   1
	//if value is less than start
	if value < dl.start.value {
		temp := dl.start //temp = 3
		dl.start = node  //dl.start = 1
		dl.start.next = temp
		nodePrev := dl.start
		dl.start.next.prev = nodePrev
		dl.start.prev = nil
		return
	}

	//if value lies in between of list
	temp := dl.start
	for temp.next != nil {
		// 3->4->8 //5
		if value < temp.next.value {
			node.next = temp.next //node.next = 8
			temp.next.prev = node //nodeNext.prev = 5
			node.prev = temp
			temp.next = node
			return
		}
		temp = temp.next
	}
}
