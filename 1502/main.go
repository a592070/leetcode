package main

import "fmt"

func main() {
	rs := canMakeArithmeticProgression([]int{
		0, 0, 0, 0,
	})

	fmt.Println(rs)
}

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) == 2 {
		return true
	}
	size := len(arr)

	min := arr[0]
	max := arr[0]
	sum := 0
	for i := 0; i < size; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
		sum += arr[i]
	}

	diff := (max - min) / (size - 1)
	for i := 1; i < size; i++ {
		tmp := (arr[i] - arr[i-1])
		if tmp != diff && tmp%diff != 0 {
			return false
		}
	}

	// shum should equal (a0 + aN) * n / 2
	rs := (min + max) * size
	if rs%2 != 0 {
		return false
	}
	return sum == rs/2
}
