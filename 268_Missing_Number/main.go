package main

import "sort"

/*
*
https://leetcode.com/problems/missing-number/
Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

Example 1:
Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.

Example 2:
Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.

Example 3:
Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.

9 -> 1
6 -> 5 -> 3 -> 4 -> 2
7
0
*/
func main() {

}

func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		if i != num {
			return i
		}
	}
	return len(nums)
}

func missingNumber2(nums []int) int {
	size := len(nums)
	actualSum := (size) * (size + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return actualSum - sum
}

func missingNumber3(nums []int) (sum int) {
	/*
		[3,0,1]
		i=0, sum = (0^3)^0 = 3^0 = 3
		i=1, sum = (1^0)^3 = 1^3 = 2
		i=2, sum = (2^1)^2 = 3^2 = 1
		result = 1 ^ 3 = 2
	*/
	for i, num := range nums {
		sum ^= i ^ num
	}
	return sum ^ len(nums)
}
