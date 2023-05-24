package main

/*
https://leetcode.com/problems/backspace-string-compare/
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.
Note that after backspacing an empty text, the text will continue empty.

Example 1:

Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Example 2:

Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

Example 3:

Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".
*/
func main() {

}
func backspaceCompare(s string, t string) bool {
	sRes := ""
	tRes := ""
	for _, char := range s {
		if char == '#' {
			if len(sRes) <= 0 {
				continue
			}
			sRes = sRes[:len(sRes)-1]
		} else {
			sRes += string(char)
		}
	}

	for _, char := range t {
		if char == '#' {
			if len(tRes) <= 0 {
				continue
			}
			tRes = tRes[:len(tRes)-1]
		} else {
			tRes += string(char)
		}
	}
	return sRes == tRes
}
