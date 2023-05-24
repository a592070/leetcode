package main

/*
https://leetcode.com/problems/maximum-depth-of-binary-tree/

Given the root of a binary tree, return its maximum depth.
A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:

Input: root = [1,null,2]
Output: 2
*/
func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth2(root *TreeNode) int {
	var subMaxDepth func(*TreeNode) int
	subMaxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := subMaxDepth(node.Left)
		right := subMaxDepth(node.Right)

		max := left
		if right > left {
			max = right
		}

		return max + 1
	}

	return subMaxDepth(root)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	max := left
	if right > left {
		max = right
	}
	return max + 1
}
