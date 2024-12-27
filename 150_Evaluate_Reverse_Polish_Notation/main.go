package main

import (
	"fmt"
	"strconv"
)

func main() {
	rs := evalRPN([]string{"2", "1", "+", "3", "*"})
	fmt.Println(rs)
}

func evalRPN(tokens []string) int {
	if len(tokens) <= 0 {
		return 0
	} else if len(tokens) == 1 {
		return toInt(tokens[0])
	}
	idx := findOperatorIndex(tokens)
	result := operator(toInt(tokens[idx-2]), toInt(tokens[idx-1]), tokens[idx])

	newTokens := append(tokens[:idx-2], toString(result))
	newTokens = append(newTokens, tokens[idx+1:]...)

	return evalRPN(newTokens)
}

func toInt(s string) int {
	atoi, _ := strconv.Atoi(s)
	return atoi
}
func toString(i int) string {
	return strconv.Itoa(i)
}

func findOperatorIndex(tokens []string) int {
	for i, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			return i
		}
	}
	return -1
}
func operator(n1, n2 int, op string) int {
	switch op {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	}
	return -1
}
