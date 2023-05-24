package main

import (
	"fmt"
	"math"
)

/*
https://leetcode.com/problems/climbing-stairs/

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Input: n = 4
Output: 5
Explanation: There are three ways to climb to the top.
1 step + 1 step + 1 step + 1 step
2 steps + 1 step + 1 step
1 step + 2 steps + 1 step
1 steps + 1 step + 2 step
2 steps + 2 step

Input: n = 5
Output: 8
Explanation: There are three ways to climb to the top.
1 step + 1 step + 1 step + 1 step + 1 step
2 steps + 1 step + 1 step + 1 step
1 step + 2 steps + 1 step + 1 step
1 steps + 1 step + 2 step + 1 step
1 steps + 1 step + 1 step + 2 step
2 steps + 2 step + 1 step
2 steps + 1 step + 2 step
1 step + 2 step + 2 step

f(n) = f(n-1) + f(n-2)
n + (n-1) + (n-2) + ... + (n-i) + ... 1
*/
func main() {
	fmt.Println(fibonacci(35))
	fmt.Println(fibonacci3(35))
}

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci2(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	res := 0
	for i := 0; i < n; i++ {
		res += factorial(n-i, i) / factorial(i, 0)
	}
	return res + 1
}

func fibonacci3(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	sqrt5 := math.Sqrt(5)

	res := (1/sqrt5)*math.Pow((1+sqrt5)/2, float64(n+1)) - (1/sqrt5)*math.Pow((1-sqrt5)/2, float64(n+1))

	return int(math.Round(res))
}

func factorial(n int, end int) int {
	if n == 0 || end == 0 {
		return 1
	}
	res := n
	for i := n - 1; i < end; i++ {
		res *= i
	}
	return res
}
