package main

/*
https://leetcode.com/problems/single-number/

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.

Example 1:

Input: nums = [2,2,1]
Output: 1

Example 2:

Input: nums = [4,1,2,1,2]
Output: 4

Example 3:

Input: nums = [1]
Output: 1
*/
func main() {
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	temp := make(map[int]bool)

	for _, val := range nums {
		temp[val] = !temp[val]
	}
	for k, val := range temp {
		if val {
			return k
		}
	}
	return 0
}

// Hint:
// Utilize the property of XOR, A ⊕ A = 0, to cancel those elements which appeared twice.
/*
(x1 xor x2 xor x3 ...) xor n xor (x1 xor x2 xor x3 ...)
=
(x1 xor x1) xor (x2 xor x2) xor (x3 xor x3) ... n
=
0 xor 0 xor 0 xor ... n
=
n

*/
func singleNumber1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := 0

	for _, val := range nums {
		res ^= val
	}

	return res
}
