package main

import (
	"sort"
)

/*
*
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]
*/
func main() {

}

func sortedSquares(nums []int) []int {
	if nums == nil || len(nums) <= 0 {
		return nums
	}
	var rs []int
	for _, val := range nums {
		rs = append(rs, val*val)
	}
	sort.Ints(rs)
	return rs
}

func sortedSquares2(nums []int) []int {
	size := len(nums)
	rs := make([]int, size)
	left, right := 0, size-1
	for i := size - 1; i >= 0; i-- {
		leftVal := nums[left] * nums[left]
		rightVal := nums[right] * nums[right]
		if rightVal >= leftVal {
			right--
			rs[i] = rightVal
		} else {
			left++
			rs[i] = leftVal
		}
	}
	return rs
}
