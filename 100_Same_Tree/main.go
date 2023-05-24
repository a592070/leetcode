package main

import "fmt"

/**
https://leetcode.com/problems/same-tree/
Given the roots of two binary trees p and q, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Example 1:

Input: p = [1,2,3], q = [1,2,3]
Output: true


Example 2:

Input: p = [1,2], q = [1,null,2]
Output: false


Example 3:

Input: p = [1,2,1], q = [1,1,2]
Output: false


Example 4:

Input: p = [10,5,15], q = [10,5,null,null,15]
10
|	\
5 	15

---
10
|
5
\
15

Output: false

*/

func main() {
	p := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 15,
		},
	}

	q := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 15,
			},
		},
	}

	fmt.Println(isSameTree(p, q))

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	fmt.Println(p, q)
	if p == q {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
