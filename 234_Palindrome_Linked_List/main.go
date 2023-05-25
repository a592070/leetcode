package main

import "fmt"

/*
https://leetcode.com/problems/palindrome-linked-list/

Given the head of a singly linked list, return true if it is a palindrome or false otherwise.

Example 1:

Input: head = [1,2,2,1]
Output: true

Example 2:

Input: head = [1,2]
Output: false

Constraints:
The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9

Follow up: Could you do it in O(n) time and O(1) space?
*/
func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	//node.println()
	//fmt.Println("================")
	//middle := middleNode(node)
	//fmt.Printf("middle=%v\n", middle)
	//fmt.Println("================")
	//rev := reverse(middle, nil)
	//rev.println()
	//fmt.Println("================")

	fmt.Printf("isPalindrome: %v", isPalindrome(node))
}

func (node *ListNode) println() {
	p := node
	for p != nil {
		fmt.Println(p)
		p = p.Next
	}
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	middle := middleNode(head)
	rev := reverse(middle, nil)

	curr := head
	curr2 := rev
	for curr != middle {
		if curr.Val != curr2.Val {
			return false
		}
		curr = curr.Next
		curr2 = curr2.Next
	}
	return true
}

func middleNode(head *ListNode) *ListNode {
	middle := head
	current := head
	currentIndex := 0
	middleIndex := 0
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

func reverse(current *ListNode, prev *ListNode) *ListNode {
	if current == nil {
		return prev
	}
	next := current.Next
	current.Next = prev
	return reverse(next, current)
}
