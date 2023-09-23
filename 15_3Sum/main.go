package main

import (
	"sort"
)

/*
*
https://leetcode.com/problems/3sum/
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
*/
func main() {
	threeSum([]int{1, 2, 3, 4})
}
func threeSum(nums []int) [][]int {
	size := len(nums)
	rs := [][]int{}
	tripletMap := make(map[[3]int][]int)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			for k := j + 1; k < size; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triplet := [3]int{nums[i], nums[j], nums[k]}
					sort.Ints(triplet[:])
					tripletMap[triplet] = triplet[:]
				}
			}
		}
	}
	for _, triplet := range tripletMap {
		rs = append(rs, triplet)
	}

	return rs
}

func threeSumV2(nums []int) [][]int {
	size := len(nums)
	sort.Ints(nums)
	rs := [][]int{}
	for i := 0; i < size-2; i++ {
		// skip same first value
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := size - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				rs = append(rs, []int{nums[i], nums[left], nums[right]})
				right--
				// skip duplicate from right
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > 0 {
				right--
				// skip duplicate from right
				//for left < right && nums[right] == nums[right+1] {
				//	right--
				//}
			} else {
				left++
				// skip duplicate from left
				//for left < right && nums[left] == nums[left-1] {
				//	left++
				//}
			}
		}
	}
	return rs
}
