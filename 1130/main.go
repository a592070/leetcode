package main

import (
	"fmt"
	"slices"
)

func main() {
	values := mctFromLeafValues([]int{1, 2, 3, 4})
	/*
		1 2 3 4
		 2   12
		   8
		=22

		1 2 3 4

	*/
	fmt.Println(values)
	/*
		1 2 3 4 5 6 7 8
		 2   12  30  56
		   8       48
		       32
		=32+8+48+2+12+30+56
	*/

	values = mctFromLeafValues([]int{6, 2, 4})
	fmt.Println(values)
	/*
		6 2 4
		 12
		   24
		=36

		6 2 4
		   8
		  24
		=32
	*/
}

func createNodes(nodes, arr1 []int) []int {
	if len(arr1) < 2 {
		return nodes
	}

	size := len(arr1)

	leftArr := arr1[:size/2]
	rightArr := arr1[size/2:]

	max1 := slices.Max(leftArr)
	max2 := slices.Max(rightArr)
	nodes = append(nodes, max1*max2)

	//left
	nodes = createNodes(nodes, leftArr)

	//right
	nodes = createNodes(nodes, rightArr)

	return nodes
}

func mctFromLeafValues(arr []int) int {
	if len(arr) == 2 {
		return arr[1] * arr[0]
	}

	if len(arr)%2 == 0 {
		nonLeafNodes := []int{}
		nonLeafNodes = createNodes(nonLeafNodes, arr)
		fmt.Println(nonLeafNodes)
		sum := 0
		for _, v := range nonLeafNodes {
			sum += v
		}

		return sum
	}

	// left weight
	leftWeightNonLeafNodes := []int{}
	leftWeightNonLeafNodes = createNodes(leftWeightNonLeafNodes, append(arr, 0))
	leftSum := 0
	for i := range leftWeightNonLeafNodes {
		leftSum += leftWeightNonLeafNodes[i]
	}

	rightWeightNonLeafNodes := []int{}
	rightWeightNonLeafNodes = createNodes(rightWeightNonLeafNodes, append([]int{0}, arr...))
	rightSum := 0
	for i := range rightWeightNonLeafNodes {
		rightSum += rightWeightNonLeafNodes[i]
	}
	return min(leftSum, rightSum)
}
