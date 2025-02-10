package main

import "math"

func main() {

}

//Given the root of a binary tree, determine if it is a valid binary search
//tree (BST).
//
// A valid BST is defined as follows:
//
//
// The left subtree of a node contains only nodes with keys less than the
//node's key.
// The right subtree of a node contains only nodes with keys greater than the
//node's key.
// Both the left and right subtrees must also be binary search trees.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isInnerBoundary(upper int, lower int, node *TreeNode) bool {
	if node == nil {
		return true
	}
	// Check node in boundary
	if lower < node.Val && node.Val < upper {
		// Set current value as new upper boundary for left nodes, and lower val boundary for right nodes.
		return isInnerBoundary(node.Val, lower, node.Left) && isInnerBoundary(upper, node.Val, node.Right)
	}
	return false
}

// isValidBST would traverse each node to check whether it's in boundary
func isValidBST(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	upperBound := math.MaxInt
	lowerBound := math.MinInt

	return isInnerBoundary(upperBound, lowerBound, root)
}
