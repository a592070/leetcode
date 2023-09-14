package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
https://leetcode.com/problems/add-binary/

Example 1:

Input: a = "11", b = "1"
Output: "100"

Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/
func main() {
	//a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	//b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	a := "11"
	b := "1"
	fmt.Println(addBinary2(a, b))
}

func addBinary(a string, b string) string {
	aInt, _ := strconv.ParseInt(a, 2, 0)
	bInt, _ := strconv.ParseInt(b, 2, 0)
	fmt.Println(aInt)
	fmt.Println(bInt)
	sum := aInt + bInt
	fmt.Println(sum)
	return strconv.FormatInt(sum, 2)
}

func addBinary2(a string, b string) string {
	aSize := len(a)
	bSize := len(b)

	max := int(math.Max(float64(aSize), float64(bSize)))
	list := []int{}
	plus := 0
	for i := 0; i < max; i++ {
		aCurr := 0
		aIdx := aSize - i - 1
		if aIdx >= 0 && a[aIdx] == '1' {
			aCurr = 1
		}

		bCurr := 0
		bIdx := bSize - i - 1
		if bIdx >= 0 && b[bIdx] == '1' {
			bCurr = 1
		}

		curr := aCurr + bCurr + plus
		if curr == 3 {
			list = append(list, 1)
			plus = 1
		} else if curr == 2 {
			list = append(list, 0)
			plus = 1
		} else if curr == 1 {
			list = append(list, 1)
			plus = 0
		} else {
			list = append(list, 0)
			plus = 0
		}

	}
	if plus == 1 {
		list = append(list, 1)
	}
	var sumSb strings.Builder
	listSize := len(list)
	for i := listSize - 1; i >= 0; i-- {
		sumSb.WriteString(strconv.Itoa(list[i]))
	}

	return sumSb.String()
}
