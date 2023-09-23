package main

/*
*
https://leetcode.com/problems/binary-tree-level-order-traversal/

Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:

Input: root = [1]
Output: [[1]]

Example 3:

Input: root = []
Output: []
*/
func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	m := make(map[int][]int)
	levelMap(root, 0, m)

	rs := make([][]int, len(m))
	for k, v := range m {
		rs[k] = v
	}

	return rs
}
func levelMap(root *TreeNode, level int, m map[int][]int) {
	if root == nil {
		return
	}
	m[level] = append(m[level], root.Val)
	levelMap(root.Left, level+1, m)
	levelMap(root.Right, level+1, m)
}
