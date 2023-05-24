package main

import "fmt"

/*
https://leetcode.com/problems/majority-element/
Example 1:

Input: nums = [3,2,3]
Output: 3

Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2
*/
func main() {
	fmt.Println(majorityElement([]int{6, 5, 5}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{3, 2, 3, 3}))
}

func majorityElement(nums []int) int {
	countMap := map[int]int{}
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	for _, val := range nums {
		countMap[val]++
	}
	fmt.Println(countMap)
	for k, v := range countMap {
		if v > size/2 {
			return k
		}
	}
	return -1
}
