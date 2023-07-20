package main

import "fmt"

/**
Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.

Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:


Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.



*/

func main() {
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}).toString())
}

func (node *TreeNode) toString() string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("{Val=%d, Left=%s, Right=%s}", node.Val, node.Left.toString(), node.Right.toString())
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

func sortedArrayToBST(nums []int) *TreeNode {
	size := len(nums)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	mid := size / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[0:mid]),
		Right: sortedArrayToBST(nums[mid+1 : size]),
	}
}
