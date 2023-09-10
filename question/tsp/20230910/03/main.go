package main

import "fmt"

func main() {
	arr := []int32{2, 5, 3, 1}
	fmt.Println(findRedundant(arr))
}

func findRedundant(planets []int32) int32 {
	size := len(planets)
	for i := 0; i < len(planets); i++ {
		var evenSum int32 = 0
		var oddSum int32 = 0

		tempArr := []int32{}
		tempArr = append(tempArr, planets[:i]...)
		tempArr = append(tempArr, planets[i+1:size]...)
		for j := 0; j < len(tempArr); j++ {
			if j%2 == 0 {
				evenSum += planets[j]
			} else {
				oddSum += planets[j]
			}
		}
		if evenSum == oddSum {
			return int32(i + 1)
		}
	}
	return 0
}
