package main

import (
	"fmt"
	"go-data-structure/linkedlist2"
)

// 1->2->3->4->nil
func main() {
	singlyLinkList := linkedlist2.NewSinglyLinkeList()
	// singlyLinkList.Add(1)
	// singlyLinkList.Add(1)

	singlyLinkList.Add(1)
	singlyLinkList.Add(2)
	singlyLinkList.Add(3)
	singlyLinkList.Add(4)
	singlyLinkList.Add(5)
	// singlyLinkList.Add(5)
	singlyLinkList.Add(6)

	// singlyLinkList.Add(0)
	// singlyLinkList.Add(4)
	// singlyLinkList.Add(5)
	// singlyLinkList.Add(2)
	// singlyLinkList.Add(0)

	// singlyLinkList.Add(1)
	// singlyLinkList.Add(1)
	// singlyLinkList.Add(1)
	// singlyLinkList.Add(1)
	// singlyLinkList.Add(1)

	fmt.Println("Linkedlist after insertion:-")
	singlyLinkList.Display()
	// singlyLinkList.Remove(10)
	// fmt.Println("Linkedlist after removal of node 3:-")
	// singlyLinkList.Display()
	// singlyLinkList.DisplayReverse(singlyLinkList.GetFirst())
	// singlyLinkList.Display()

	/******** Problems ********/
	//1. Remove Nth Node From End of List
	//singlyLinkList.RemoveNthFromEnd(singlyLinkList.GetFirst(), 2)
	//singlyLinkList.Display()

	//2. Merge Two Sorted Lists
	// singlyLinkList2 := linkedlist2.NewSinglyLinkeList()
	// singlyLinkList2.SortedAdd(1)
	// singlyLinkList2.SortedAdd(3)
	// singlyLinkList2.SortedAdd(4)
	// singlyLinkList2.MergeTwoLists(singlyLinkList.GetFirst(), singlyLinkList2.GetFirst())
	// singlyLinkList.Display()

	//3. Swap Nodes in Pairs
	// singlyLinkList.SwapPairs(singlyLinkList.GetFirst())
	// singlyLinkList.Display()

	//4. Rotate linked list
	// singlyLinkList.RotateRight(singlyLinkList.GetFirst(), 10)
	// singlyLinkList.Display()

	//5. Remove Duplicates from sorted list
	// singlyLinkList.DeleteDuplicates(singlyLinkList.GetFirst())
	// singlyLinkList.Display()

	//6. Remove Duplicates from sorted listII
	// singlyLinkList.DeleteDuplicates2(singlyLinkList.GetFirst())
	// singlyLinkList.Display()

	//7. Remove Duplicates from sorted listII
	// singlyLinkList.ReverseBetween(singlyLinkList.GetFirst(), 2, 4)
	// singlyLinkList.Display()

	//8. Swap Nodes in Pairs
	// singlyLinkList.SwapPairs(singlyLinkList.GetFirst())
	// singlyLinkList.Display()

	//9. Partition List
	// singlyLinkList.Partition(singlyLinkList.GetFirst(), 0)
	// singlyLinkList.Display()

	//10.  Swapping Nodes in a Linked List
	// singlyLinkList.SwapNodes(singlyLinkList.GetFirst(), 5)
	// singlyLinkList.Display()

	//11. Merge In Between Linked Lists
	// singlyLinkList2 := linkedlist2.NewSinglyLinkeList()
	// singlyLinkList2.Add(1000000)
	// singlyLinkList2.Add(1000001)
	// singlyLinkList2.Add(1000002)
	// singlyLinkList2.Add(1000003)
	// singlyLinkList2.Add(1000004)
	// singlyLinkList.MergeInBetween(singlyLinkList.GetFirst(), 2, 5, singlyLinkList2.GetFirst())
	// singlyLinkList.Display()

	//12. Delete middle node of linked list
	// singlyLinkList.DeleteMiddle(singlyLinkList.GetFirst())
	// singlyLinkList.Display()

	//13. Remove Nodes From Linked List
	// singlyLinkList.RemoveNodes(singlyLinkList.GetFirst())
	// singlyLinkList.Display()

	//14. Merge Nodes in Between Zeros
	// singlyLinkList.MergeNodes(singlyLinkList.GetFirst())
	// singlyLinkList.Display()

	//15. Reverse linked list
	// singlyLinkList.ReverseList(singlyLinkList.GetFirst())
	// singlyLinkList.Display()

	//16. Palindrome Linked List
	// val := singlyLinkList.IsPalindrome(singlyLinkList.GetFirst())
	// fmt.Println("is palindrome is>>>", val)
	// singlyLinkList.Display()

	//17. Odd Even Linked List
	// singlyLinkList.OddEvenList(singlyLinkList.GetFirst())
	// singlyLinkList.Display()

	//18. Remove elements from Linked List
	// singlyLinkList.RemoveElements(singlyLinkList.GetFirst(), 7)
	// singlyLinkList.Display()

	//19.  Middle of the Linked List
	singlyLinkList.MiddleNode(singlyLinkList.GetFirst())
	singlyLinkList.Display()
}
