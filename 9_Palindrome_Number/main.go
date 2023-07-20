package main

/*
Given an integer x, return true if x is a palindrome, and false otherwise.


Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.


Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.


Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Follow up: Could you solve it without converting the integer to a string?
*/

func main() {

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var arr []int
	for x > 0 {
		arr = append([]int{x % 10}, arr...)
		x /= 10
	}
	size := len(arr)
	for i := range arr {
		if arr[i] != arr[size-1-i] {
			return false
		}
	}
	return true
}
