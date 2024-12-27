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

func isOperator(operator string) bool {
	switch operator {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	}
	return false
}

type stack struct {
	collection []int
}

func (s *stack) pop() int {
	ele := s.collection[len(s.collection)-1]
	s.collection = s.collection[:len(s.collection)-1]
	return ele
}
func (s *stack) push(element int) {
	s.collection = append(s.collection, element)
}
func (s *stack) len() int {
	return len(s.collection)
}

func evalRPNUsingStack(tokens []string) int {
	if len(tokens) <= 0 {
		return 0
	} else if len(tokens) == 1 {
		return toInt(tokens[0])
	}

	s := stack{[]int{}}
	for _, token := range tokens {
		if isOperator(token) {
			n1 := s.pop()
			n2 := s.pop()
			// should reverse the number's order
			rs := operator(n2, n1, token)
			s.push(rs)
		} else {
			s.push(toInt(token))
		}
	}

	return s.pop()
}
