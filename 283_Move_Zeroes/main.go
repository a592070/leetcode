package main

import "fmt"

/*
https://leetcode.com/problems/move-zeroes/

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:

Input: nums = [0]
Output: [0]
*/
func main() {
	nums := []int{0, 1, 0, 3, 12}
	//nums := []int{0, 0, 0, 0}
	//nums := []int{0, 0, 0, 1}
	//nums = removeAtIndex(nums, 2)
	//fmt.Println(nums)
	//for i, val := range nums {
	//	if val == 0 {
	//		nums = removeAtIndex(nums, i)
	//	}
	//}
	moveZeroes(nums)
	fmt.Println(nums)
}

func removeAtIndex(arr []int, index int) []int {
	if index+1 >= len(arr) {
		return append(arr[:index])
	}
	return append(arr[:index], arr[index+1:]...)
}

func moveZeroes(nums []int) {
	size := len(nums)
	if size == 1 {
		return
	}

	for i := 0; i < size; i++ {
		if nums[i] == 0 {
			for j := i; j < size; j++ {
				if nums[j] != 0 {
					temp := nums[i]
					nums[i] = nums[j]
					nums[j] = temp
					break
				}
			}
		}
	}
}

func moveZeroes1(nums []int) {
	size := len(nums)
	if size == 1 {
		return
	}

	noneZeroIndex := 0
	for i, val := range nums {
		if val != 0 {
			nums[i], nums[noneZeroIndex] = nums[noneZeroIndex], nums[i]
			noneZeroIndex++
		}
	}
}
