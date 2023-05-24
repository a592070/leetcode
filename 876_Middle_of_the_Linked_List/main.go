package main

/*
https://leetcode.com/problems/middle-of-the-linked-list/
*/
func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	middle := head
	current := head
	currentIndex := 0
	middleIndex := 0
	// 1
	// 1,2
	// 1,2,3
	for current.Next != nil {
		currentIndex++
		if currentIndex > middleIndex*2 {
			middle = middle.Next
			middleIndex++
		}
		current = current.Next
	}
	return middle
}

func middleNode2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast := head
	slow := head

	for fast != nil {
		if fast.Next == nil {
			if slow.Next != nil {
				slow = slow.Next
			}
			break
		}

		fast = fast.Next.Next
		slow = slow.Next

	}
	return slow
}
