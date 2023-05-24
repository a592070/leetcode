package main

import "fmt"

/*
*
https://leetcode.com/problems/reverse-linked-list/

Example 1:

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:

Input: head = [1,2]
Output: [2,1]

Example 3:

Input: head = []
Output: []
*/
func main() {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	curr := head
	for i := 1; i < 10; i++ {
		curr.Next = &ListNode{
			Val: i,
		}
		curr = curr.Next
	}

	print(reverse2(head))
}

func reverse(current *ListNode, prev *ListNode) *ListNode {
	if current == nil {
		return prev
	}
	next := current.Next
	current.Next = prev
	return reverse(next, current)
}

func reverse2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	var prev *ListNode
	//prev := &ListNode{}
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func print(node *ListNode) {
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
