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

func (sl *SinglyLinkeList) GetFirst() *Node {
	return sl.start
}

func (sl *SinglyLinkeList) Add(value int) {
	node := newNode(value)
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

func (sl *SinglyLinkeList) DisplayReverse(head *Node) {
	if head == nil {
		fmt.Println("nil")
		return
	}

	sl.DisplayReverse(head.next)
	fmt.Printf("%d->", head.value)
	return
}

func (sl *SinglyLinkeList) CountNodes(head *Node) int {
	count := 0
	current := head
	for current != nil {
		current = current.next
		count += 1
	}
	return count
}

// 65
func (sl *SinglyLinkeList) RemoveNthFromEnd(head *Node, n int) *Node {
	//count all nodes
	count := sl.CountNodes(head)

	//goto previous position of deleting node
	removePosition := count - n
	//if deleting node is 1
	if removePosition == 0 {
		sl.start = head.next
		fmt.Println("removePosition, count>>>>", removePosition, count)
		return head.next
	}
	temp := head
	for i := 1; i < removePosition; i++ {
		temp = temp.next
	}

	//changing link for deleting node
	temp.next = temp.next.next
	return head
}

func (sl *SinglyLinkeList) MergeTwoLists(list1 *Node, list2 *Node) *Node {
	//base case if list1 is empty return list2
	if list1 == nil {
		return list2
	}
	//base case if list2 is empty return list1
	if list2 == nil {
		return list1
	}
	//create start empty node new merged linkedlist, temp variable for iteration of linked list
	var start, temp *Node
	for list1 != nil && list2 != nil {
		//temp2 variable for creating new node value
		var temp2 *Node
		if list1.value < list2.value {
			temp2 = list1
			list1 = list1.next
			temp2.next = nil
		} else {
			temp2 = list2
			list2 = list2.next
			temp2.next = nil
		}

		if temp == nil {
			//if temp is empty, assign first node of temp2 to start & temp,
			//after that will update the temp only, so that start node will access the same address of temp
			start = temp2
			temp = temp2
		} else {
			temp.next = temp2
			temp = temp.next
		}
	}
	if list1 != nil {
		temp.next = list1
	} else {
		temp.next = list2
	}
	return start
}

func (sl *SinglyLinkeList) SwapPairs(head *Node) *Node {
	//base case : if linkedlist is empty or have single node then return node  itself
	if head == nil || head.next == nil {
		return head
	}
	//firstly swap first two nodes
	first := head
	head = head.next
	first.next = head.next
	head.next = first
	prev := head.next
	//fmt.Println("head & head.next......", head.value, head.next.value, prev.value)

	for prev.next != nil && prev.next.next != nil {
		first := prev.next
		second := prev.next.next

		prev.next = second
		first.next = second.next
		second.next = first
		prev = first
	}
	return head
}

func (sl *SinglyLinkeList) RotateRight(head *Node, k int) *Node {
	count := sl.CountNodes(head)
	if head == nil || head.next == nil || k == 0 {
		return head
	}
	k = k % count
	if k == 0 {
		return head
	}
	startPosition := count - k
	current1 := head
	var newNode *Node
	for i := 1; i < startPosition; i++ {
		current1 = current1.next
	}
	current2 := current1.next
	current1.next = nil

	for i := startPosition + 1; i <= count; i++ {
		if newNode == nil {
			newNode = current2
		}
		if current2.next != nil {
			current2 = current2.next
		}
	}
	current2.next = head
	return newNode
}

func (sl *SinglyLinkeList) DeleteDuplicates(head *Node) *Node {
	//base case
	if head == nil || head.next == nil {
		return head
	}
	start := head
	//1->1->2->2->nil
	for start.next != nil {
		if start.value == start.next.value {
			start.next = start.next.next
		} else {
			start = start.next
		}
	}
	return head
}

func (sl *SinglyLinkeList) DeleteDuplicates2(head *Node) *Node {
	//basecase
	if head == nil || head.next == nil {
		return head
	}

	start := head
	//var myapp map[int]int
	newMap := make(map[int]int)
	for start != nil {
		_, isPresent := newMap[start.value]
		if isPresent {
			newMap[start.value] += 1
		} else {
			newMap[start.value] = 1
		}
		start = start.next
	}

	for i := 0; i < len(newMap); i++ {

	}
	fmt.Println("here MAPPPPPPPP1", newMap)
	return head
}

func (sl *SinglyLinkeList) ReverseMap(head *Node, left int, right int) map[int]int {
	var myapp map[int]int
	myapp = make(map[int]int)

	start := head
	i := 1
	for start != nil {
		myapp[i] = start.value
		start = start.next
		i = i + 1
	}

	for left < right {
		temp := myapp[left]
		myapp[left] = myapp[right]
		myapp[right] = temp
		left++
		right--
	}
	return myapp
}

func (sl *SinglyLinkeList) ReverseBetween(head *Node, left int, right int) *Node {
	//basecase
	if head == nil || head.next == nil {
		return head
	}

	//create map for storing key/values of linked list & reverse it
	map1 := sl.ReverseMap(head, left, right)

	start := head
	//change the value of linked list from map
	for i := 1; i <= len(map1); i++ {
		start.value = map1[i]
		start = start.next
	}
	return head
}

func (sl *SinglyLinkeList) Partition(head *Node, x int) *Node {
	if head == nil || head.next == nil {
		return head
	}
	var lessNodes, temp, greatedNodes, temp2 *Node
	start := head
	for start != nil {
		if start.value < x {
			if lessNodes == nil {
				lessNodes = start
				temp = start
			} else {
				temp.next = start
				temp = temp.next
			}
		}

		if start.value >= x {
			if greatedNodes == nil {
				greatedNodes = start
				temp2 = start
			} else {
				temp2.next = start
				temp2 = temp2.next
			}
		}
		start = start.next
	}
	if temp == nil && temp2 != nil {
		return greatedNodes
	} else if temp2 == nil && temp != nil {
		return lessNodes
	}
	temp.next = greatedNodes
	temp2.next = nil
	return lessNodes
}

func (sl *SinglyLinkeList) GetKthNode(head *Node, count, k int) *Node {
	start := head
	for i := 1; i <= count; i++ {
		if i == k {
			return start
		}
		start = start.next
	}
	return start
}

func (sl *SinglyLinkeList) SwapNodes(head *Node, k int) *Node {
	//basecase
	if head == nil || head.next == nil || k == 0 {
		return head
	}
	//get length of linkedlist
	count := sl.CountNodes(head)
	getFirstNode := sl.GetKthNode(head, count, k)
	getLastNode := sl.GetKthNode(head, count, count-(k-1))

	//swap 2 nodes
	tempValue := getFirstNode.value
	temp2 := getLastNode

	getFirstNode.value = temp2.value
	getLastNode.value = tempValue
	return head
}

func (sl *SinglyLinkeList) GetKthPositionNode(head *Node, count, k int) *Node {
	start := head
	for i := 0; i < count; i++ {
		if i == k {
			return start
		}
		start = start.next
	}
	return start
}

func (sl *SinglyLinkeList) MergeInBetween(list1 *Node, a int, b int, list2 *Node) *Node {
	//get length of linkedlist
	count := sl.CountNodes(list1)
	//get previous node from which we have to delete
	prevNode := sl.GetKthPositionNode(list1, count, a-1)

	//get next node to which we have link list2
	nextNode := sl.GetKthPositionNode(list1, count, b+1)
	prevNode.next = list2
	for list2.next != nil {
		list2 = list2.next
	}
	//fmt.Println("list2.valueeeee.", list2.value)
	list2.next = nextNode
	return list1
}

// 1->2->3->4
func (sl *SinglyLinkeList) DeleteMiddle(head *Node) *Node {
	//get length of linked list
	length := sl.CountNodes(head)
	if length == 1 {
		return nil
	}
	middle := length / 2
	start := head

	fmt.Println("middleeee.", middle, length)
	for i := 0; i < length; i++ {
		if start.next != nil {
			if i+1 == middle && start.next.next != nil {
				start.next = start.next.next
			} else if i+1 == middle && start.next.next == nil {
				start.next = nil
				break
			}
			start = start.next
		}
	}
	return head
}

func (sl *SinglyLinkeList) RemoveNodes(head *Node) *Node {
	//basecase
	if head == nil || head.next == nil {
		return head
	}
	start := head
	var newlist, temp *Node

	if head.value == head.next.value {
		newlist = start
		temp = start
	}
	for start.next != nil {
		if start.next.value >= start.value {
			if newlist == nil {
				newlist = start.next
				temp = start.next
			} else {
				fmt.Println("temp values.", temp.value)
				temp.next = start.next
				temp = temp.next
			}
		}
		start = start.next
	}

	for newlist != nil {
		fmt.Println("Newlist.", newlist.value)
		newlist = newlist.next
	}
	// for temp != nil {
	// 	fmt.Println("Newlist.", temp.value)
	// 	temp = temp.next
	// }
	return temp
}

func (sl *SinglyLinkeList) MergeNodes(head *Node) *Node {
	start := head.next
	if head == nil || head.next == nil {
		return head
	}
	total := 0
	var newlist, temp *Node
	for start != nil {
		if start.value != 0 {
			total += start.value
		} else {
			newNode := &Node{
				value: total,
			}
			if newlist == nil {
				newlist = newNode
				temp = newNode
			} else {
				temp.next = newNode
				temp = newNode
			}
			total = 0
		}
		start = start.next
	}
	return newlist
}

func (sl *SinglyLinkeList) ReverseList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	var stack []int
	start := head
	for start != nil {
		stack = append(stack, start.value)
		start = start.next
	}
	top := len(stack) - 1
	for head != nil && top >= 0 {
		head.value = stack[top]
		head = head.next
		top--
	}
	return head
}

func (sl *SinglyLinkeList) IsPalindrome(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}
	var stack []int
	start := head
	for start != nil {
		stack = append(stack, start.value)
		start = start.next
	}
	n := len(stack) - 1
	i := 0
	for i < n {
		if stack[i] != stack[n] {
			return false
		}
		i++
		n--
	}
	return true
}

func (sl *SinglyLinkeList) OddEvenList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	var oddList, temp1, evenList, temp2 *Node
	if head.next != nil {
		//Assign first node of oddList
		oddList = head
		temp1 = head

		//Assign first node of evenList
		evenList = head.next
		temp2 = head.next
	}
	i := 1
	odd := true
	start := head.next.next
	for start != nil {
		if odd == true {
			temp1.next = start
			temp1 = start
			odd = false
		} else {
			temp2.next = start
			temp2 = start
			odd = true
		}
		i += 1
		start = start.next
	}
	temp2.next = nil
	temp1.next = evenList
	return oddList
}

func (sl *SinglyLinkeList) RemoveElements(head *Node, val int) *Node {
	if head == nil {
		return nil
	}
	var newHead, temp *Node
	start := head
	for start != nil {
		if start.value != val {
			if newHead == nil {
				newHead = start
				temp = start
			} else {
				temp.next = start
				temp = start
			}
		}
		start = start.next
	}
	if temp != nil {
		temp.next = nil
	}
	return newHead
}

func (sl *SinglyLinkeList) MiddleNode(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	//get length of linked list
	length := sl.CountNodes(head)
	middle := length / 2

	getNode := sl.GetKthNode(head, length, middle+1)
	for getNode != nil {
		getNode = getNode.next
	}
	return getNode
}
