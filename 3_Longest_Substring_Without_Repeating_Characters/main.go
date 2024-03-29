package main

/**
https://leetcode.com/problems/longest-substring-without-repeating-characters/
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.


Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.


Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

*/

func main() {

}

func lengthOfLongestSubstring(s string) int {
	dict := [256]int{}
	maxLen := 0
	start := 0
	for i := 0; i < len(s); i++ {

		if start < dict[s[i]] {
			start = dict[s[i]]
		}
		if maxLen < i-start+1 {
			maxLen = i - start + 1
		}
		dict[s[i]] = i + 1
	}
	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	result := 0
	start := 0
	charIndexMap := make(map[uint8]int)
	for end := 0; end < len(s); end++ {
		duplicateIndex, isDuplicate := charIndexMap[s[end]]
		if isDuplicate {
			result = max(result, end-start)

			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}
			start = duplicateIndex + 1
		}
		charIndexMap[s[end]] = end
	}
	result = max(result, len(s)-start)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
