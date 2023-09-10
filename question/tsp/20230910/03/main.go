package main

import "fmt"

func main() {
	//arr := []int32{2, 5, 3, 1}
	arr := []int32{2, 3, 1, 4, 4, 10, 5, 5}
	fmt.Println(findRedundant(arr))
}

/*
*
Let even sum equals odd sum when remove the redundant number.
*/
func findRedundant(planets []int32) int {
	size := len(planets)
	for i := 0; i < len(planets); i++ {
		var evenSum int32 = 0
		var oddSum int32 = 0

		tempArr := []int32{}
		tempArr = append(tempArr, planets[:i]...)
		tempArr = append(tempArr, planets[i+1:size]...)
		for j := 0; j < len(tempArr); j++ {
			if j%2 == 0 {
				evenSum += tempArr[j]
			} else {
				oddSum += tempArr[j]
			}
		}
		if evenSum == oddSum {
			return i
		}
	}
	return 0
}
