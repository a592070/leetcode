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

var res = true

func main() {
	var check func(node *TreeNode) int
	check = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := check(node.Left)
		right := check(node.Right)
		if left == -1 || right == -1 || math.Abs(float64(left-right)) > 1 {
			return -1
		}
		return int(math.Max(float64(left), float64(right)) + 1)
	}
	isBalanced := check(nil) != -1
	fmt.Println(isBalanced)
}

func maxDep(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if !res {
		return 0
	}
	leftDep := maxDep(node.Left)
	rightDep := maxDep(node.Right)
	if math.Abs(float64(leftDep-rightDep)) > 1 {
		res = false
		return 0
	}
	return int(math.Max(float64(leftDep), float64(rightDep)) + 1)
}
