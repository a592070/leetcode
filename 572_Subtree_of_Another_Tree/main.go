package main

/*
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.
A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.

Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true

Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
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

func equals(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil || node2 == nil {
		if node1 == node2 {
			return true
		}
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	return equals(node1.Left, node2.Left) && equals(node1.Right, node2.Right)
}
func (node1 *TreeNode) equals(node2 *TreeNode) bool {
	if node1 == nil || node2 == nil {
		if node1 == node2 {
			return true
		}
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	return node1.Left.equals(node2.Left) && node1.Right.equals(node2.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	/*
		wrong part
		consider root=[1,1], subTree=[1]
	*/
	//find := findSubTree(root, subRoot.Val)
	//return equals(subRoot, find)

	var arr []*TreeNode
	arr = findAllSubTree(arr, root, subRoot.Val)
	for _, node := range arr {
		if equals(subRoot, node) {
			return true
		}
	}
	return false
}
func findSubTree(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	rs := findSubTree(root.Left, val)
	if rs != nil {
		return rs
	}
	rs = findSubTree(root.Right, val)
	return rs
}

func findAllSubTree(arr []*TreeNode, root *TreeNode, val int) []*TreeNode {
	if root == nil {
		return arr
	}
	if root.Val == val {
		arr = append(arr, root)
	}
	arr = findAllSubTree(arr, root.Left, val)
	arr = findAllSubTree(arr, root.Right, val)
	return arr
}
