package main

import (
	"fmt"
	"strings"
)

/*
https://leetcode.com/problems/roman-to-integer/

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.

Example 1:

Input: s = "III"
Output: 3
Explanation: III = 3.

Example 2:

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

Example 3:

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

Example 3:

Input: s = "IMICIXIX"
Output: 3999
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/
func main() {

	fmt.Println(intToRoman(3999))
	fmt.Println(romanToInt(intToRoman(3999)))
}

func romanToInt(s string) int {
	constant := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	res := 0

	sArr := []rune(s)
	size := len(sArr)

	for i := 0; i < size; i++ {
		if i+1 < size && constant[sArr[i]] < constant[sArr[i+1]] {
			res -= constant[sArr[i]]
		} else {
			res += constant[sArr[i]]
		}
	}
	return res
}

func intToRoman(num int) string {
	romanArr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	intArr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	res := ""

	for i, k := range romanArr {
		v := intArr[i]
		if num >= v {
			res += strings.Repeat(k, num/v)
		}
		num %= v
	}
	return res
}
