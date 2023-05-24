package main

/*
https://leetcode.com/problems/counting-bits/
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

Example 1:

Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10

Example 2:

Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101
*/
func main() {

}

func countBits(n int) []int {
	res := []int{0}
	if n == 0 {
		return res
	}
	if n == 1 {
		return []int{0, 1}
	}
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			// because the last digit is 0
			res = append(res, res[i>>1])
		} else {
			res = append(res, res[i-1]+1)
		}
	}
	return res
}

/*
0 0
1 1
2 10
3 11
4 100
5 101
6 110
7 111
8 1000
9 1001
10 1010
11 1011
12 1100
13 1101
14 1110
15 1111
16 10000
*/
