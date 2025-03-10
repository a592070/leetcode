package main

import "fmt"

func main() {
	fmt.Println(bubbleSort([]int{5, 4, 3, 2, 1}))
	fmt.Println(bubbleSort([]int{5, 1, 3, 2, 4}))
}

func bubbleSort(arr []int) []int {
	size := len(arr)
	if size <= 1 {
		return arr
	} else if size == 2 {
		arr[0], arr[1] = arr[1], arr[0]
		return arr
	}

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
