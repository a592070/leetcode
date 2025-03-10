package main

import "fmt"

func main() {
	orderOfLargestPlusSign(1, [][]int{{0, 0}})
	orderOfLargestPlusSign(5, [][]int{{4, 2}})
	/*
		00000
		00000
		00001
		00000
		00000
	*/
	orderOfLargestPlusSign(2, [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}})
	orderOfLargestPlusSign(3, [][]int{{0, 0}, {1, 2}})
	// 1 0 0
	// 0 0 0
	// 0 1 0

	/*
			0 1 2
			1 2 3
			1 0 1

			0 1 1
			1 2 2
			2 0 3

		    0 2 1
			3 2 1
			1 0 1

			0 2 3
			2 1 2
			1 0 1
	*/
	orderOfLargestPlusSign(2, [][]int{{0, 0}, {0, 1}, {1, 0}})
	/*
		11
		10
	*/
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	maps := make([][]int, n)

	right := make([][]int, n)
	down := make([][]int, n)
	up := make([][]int, n)
	left := make([][]int, n)

	for i := 0; i < n; i++ {
		maps[i] = make([]int, n)
		right[i] = make([]int, n)
		down[i] = make([]int, n)
		up[i] = make([]int, n)
		left[i] = make([]int, n)
		for j := 0; j < n; j++ {
			maps[i][j] = 1
			right[i][j] = 1
			down[i][j] = 1
			up[i][j] = 1
			left[i][j] = 1
		}
	}
	//fmt.Println(maps)
	for _, mine := range mines {
		x := mine[0]
		y := mine[1]
		maps[y][x] = 0
		right[y][x] = 0
		down[y][x] = 0
		up[y][x] = 0
		left[y][x] = 0
	}
	//fmt.Println(maps)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if maps[i][j] == 1 {
				if i > 0 {
					down[i][j] = down[i-1][j] + 1
				}
				if j > 0 {
					right[i][j] = right[i][j-1] + 1
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if maps[i][j] == 1 {
				if i < n-1 {
					up[i][j] = up[i+1][j] + 1
				}
				if j < n-1 {
					left[i][j] = left[i][j+1] + 1
				}
			}
		}
	}

	//fmt.Println(right)
	//fmt.Println(down)
	//fmt.Println(left)
	//fmt.Println(up)

	rs := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rs = max(rs, min(right[i][j], down[i][j], left[i][j], up[i][j]))
		}
	}
	fmt.Println(rs)
	return rs
}
