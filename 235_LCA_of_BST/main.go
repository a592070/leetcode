package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 9},
		},
	}

	//fmt.Println(findIndex(root, 9, 0))
	fmt.Println(lowestCommonAncestor(root, &TreeNode{Val: 7}, &TreeNode{Val: 2}))

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	rootVal := root.Val
	pVal := p.Val
	qVal := q.Val
	if rootVal == pVal || rootVal == qVal {
		return root
	}

	min := int(math.Min(float64(pVal), float64(qVal)))
	max := int(math.Max(float64(pVal), float64(qVal)))
	if max < rootVal {
		return lowestCommonAncestor(root.Left, p, q)
	} else if min > rootVal {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

func findIndex(root *TreeNode, target int, offset int) int {
	if root.Val == target {
		return offset
	}
	if target < root.Val {
		return findIndex(root.Left, target, offset+1)
	}
	offset++
	if target > root.Val {
		return findIndex(root.Right, target, offset+2)
	}
	return -1
}
