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

func (sl *SinglyLinkeList) Display(head *Node) {
	var current *Node
	if head == nil {
		current = sl.start
	} else {
		current = head
	}
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
	var stack []*Node

	for start != nil {
		stack = append(stack, start)
		start = start.next
	}

	top := len(stack) - 1
	max := stack[top]
	top--
	for top >= 0 {
		if stack[top].value >= max.value {
			stack[top].next = max
			max = stack[top]
		}
		top--
	}
	return max
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

func (sl *SinglyLinkeList) GetMiddleList(head *Node) *Node {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func (sl *SinglyLinkeList) ReorderList(head *Node) {
	//Get middle of linked list
	mid := sl.GetMiddleList(head)
	start := mid.next
	mid.next = nil

	//push element into stack
	var stack []int
	for start != nil {
		stack = append(stack, start.value)
		start = start.next
	}

	current := head
	top := len(stack) - 1
	for top >= 0 {
		next := current.next
		newHead := &Node{value: stack[top]}
		newHead.next = next
		current.next = newHead
		current = next
		top--
	}
}

func (sl *SinglyLinkeList) AddTwoNumbers(l1 *Node, l2 *Node) *Node {
	var result, temp *Node
	remainder := 0

	//iterate over l1 & l2 list
	for l1 != nil || l2 != nil {
		//create new node
		newNode := &Node{
			value: 0,
		}
		//add value of l1 & l2 to newNode of value
		if l1 != nil {
			newNode.value += l1.value
			l1 = l1.next
		}
		if l2 != nil {
			newNode.value += l2.value
			l2 = l2.next
		}

		newNode.value += remainder
		remainder = 0
		if newNode.value > 9 {
			newNode.value = newNode.value - 10
			remainder = 1
		}
		if result == nil {
			temp = newNode
			result = newNode
		} else {
			temp.next = newNode
			temp = temp.next
		}
	}
	if remainder > 0 {
		newNode := &Node{
			value: remainder,
		}
		temp.next = newNode
	}
	return result
}

func (sl *SinglyLinkeList) InsertionSortList(head *Node) *Node {
	//get count of nodes in linked list
	n := sl.CountNodes(head)
	fmt.Println("Number of nodes.....", n)

	//headPrev := head
	headNext := head.next
	head.next = nil
	for headNext != nil {

	}
	return head
}

type nodeIndex struct {
	node  *Node
	index int
}

func (sl *SinglyLinkeList) NextLargerNodes(head *Node) []int {
	n := sl.CountNodes(head)
	var stack []nodeIndex
	var arr = make([]int, n)
	current := head
	i := 0
	top := -1
	for current != nil {
		for top > -1 && stack[top].node.value < current.value {
			arr[stack[top].index] = current.value
			stack = stack[:top]
			top--
		}
		stack = append(stack, nodeIndex{node: current, index: i})
		current = current.next
		i++
		top++
	}
	return arr
}

/*
type Node struct {
    val int
    next *Node
}

type MyLinkedList struct {
    head *Node
    tail *Node
}


func Constructor() MyLinkedList {
    return MyLinkedList{
        head:nil,
        tail:nil,
    }
}

func (this *MyLinkedList) print(key string){
    fmt.Print(key)
    current:=this.head
    for current!=nil{
        fmt.Printf("%d->",current.val)
        current=current.next
    }
    fmt.Println("nil")
}

func (this *MyLinkedList) Get(index int) int {

    current := this.head
    for i:=0;i<index && current!=nil;i++{
        current = current.next
    }
    if current==nil{
        return -1
    }
    return current.val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    defer this.print("AddAtHead")
    newNode := &Node{
        val : val,
        next : this.head,
    }
    if this.tail == nil {
        this.tail = newNode
    }
    this.head = newNode
}


func (this *MyLinkedList) AddAtTail(val int)  {
    defer this.print("AddAtTail")
    newEnd := &Node{
        val : val,
        next : nil,
    }

   if this.tail == nil {
       this.tail =  newEnd
       this.head = newEnd
       return
   }
   this.tail.next = newEnd
   this.tail = newEnd
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    defer this.print("AddAtIndex")
    current := this.head
    newNode := &Node{
        val : val,
        next : nil,
    }
    if index == 0 {
        newNode.next = this.head
        this.head = newNode
        if this.tail==nil{
            this.tail = newNode
        }
        return
    }
    index-=1
    for i:=0;i<index;i++{
        current=current.next
    }
    newNode.next=current.next

    current.next=newNode
    if current == this.tail{
        this.tail = newNode
    }
    return
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    defer this.print("DeleteAtIndex")
    if index == 0 {
        this.head = this.head.next
        return
    }

    current := this.head
    for index > 1 && current!=nil {
        current = current.next
        index --
    }
    if current==nil||current.next==nil{
        return
    }
    current.next = current.next.next
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
*/
