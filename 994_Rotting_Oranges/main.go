package main

func main() {
}

func orangesRotting(grid [][]int) int {
	rotted := [][]int{}
	freshCount := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				rotted = append(rotted, []int{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}
	if len(rotted) < 1 {
		return -1
	}

	rotateArround := func(grid [][]int, rotted [][]int, freshCount int, x, y int) ([][]int, [][]int, int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			return grid, rotted, freshCount
		}
		if x-1 >= 0 && grid[x-1][y] == 1 {
			grid[x-1][y] = 2
			rotted = append(rotted, []int{x - 1, y})
			freshCount--
		}
		if y-1 >= 0 && grid[x][y-1] == 1 {
			grid[x][y-1] = 2
			rotted = append(rotted, []int{x, y - 1})
			freshCount--
		}
		if x+1 < len(grid) && grid[x+1][y] == 1 {
			grid[x+1][y] = 2
			rotted = append(rotted, []int{x + 1, y})
			freshCount--
		}
		if y+1 < len(grid[0]) && grid[x][y+1] == 1 {
			grid[x][y+1] = 2
			rotted = append(rotted, []int{x, y + 1})
			freshCount--
		}
		return grid, rotted, freshCount
	}

	step := 0
	for len(rotted) > 0 {
		var nextRotted [][]int
		for i := 0; i < len(rotted); i++ {
			grid, nextRotted, freshCount = rotateArround(grid, nextRotted, freshCount, rotted[i][0], rotted[i][1])
		}
		rotted = nextRotted
		step++
		if freshCount == 0 {
			return step
		}
	}
	return -1
}
