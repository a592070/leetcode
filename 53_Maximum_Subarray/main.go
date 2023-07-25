package main

/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
*/
func main() {
}

func maxSubArray(nums []int) int {
	if nums == nil {
		return 0
	}
	size := len(nums)
	if size <= 0 {
		return 0
	}

	max := nums[0]
	sum := 0
	for i := 0; i < size; i++ {
		sum += nums[i]
		if sum > max {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}
	}
	return max
}
