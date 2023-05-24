package main

/*
https://leetcode.com/problems/longest-palindrome/

Example 1:

Input: s = "abccccdd"
Output: 7
Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.

Example 2:

Input: s = "a"
Output: 1
Explanation: The longest palindrome that can be built is "a", whose length is 1.
*/
func main() {

}

func longestPalindrome(s string) int {
	singleMap := map[uint8]bool{}
	doubleMap := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		char := s[i]
		if singleMap[char] {
			doubleMap[char] += 2
			delete(singleMap, char)
		} else {
			singleMap[char] = true
		}
	}
	count := 0

	for k := range doubleMap {
		count += doubleMap[k]
	}
	if len(singleMap) > 0 {
		count++
	}

	return count
}

func longestPalindrome2(s string) int {

	return 0
}
