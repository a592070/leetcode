package main

/*
https://leetcode.com/problems/diameter-of-binary-tree/

Example 1:

Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

Example 2:

Input: root = [1,2]
Output: 1
*/
func main() {

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

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var nextOrder func(node *TreeNode) int
	nextOrder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := nextOrder(node.Left)
		right := nextOrder(node.Right)

		// update res val
		if left+right > res {
			res = left + right
		}

		// count sub tree depth from bottom
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}

	nextOrder(root)
	return res
}
