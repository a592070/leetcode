package main

func main() {

}

func numIslands(grid [][]byte) int {
	mLen := len(grid)
	nLen := len(grid[0])
	traveled := make([][]bool, len(grid))

	for i := range traveled {
		traveled[i] = make([]bool, len(grid[i]))
	}

	var dfs func(traveled [][]bool, i, j int)
	dfs = func(traveled [][]bool, i, j int) {
		if i < 0 || i >= mLen || j < 0 || j >= nLen {
			return
		}
		if traveled[i][j] {
			return
		}
		traveled[i][j] = true
		if grid[i][j] == '0' {
			return
		} else {
			dfs(traveled, i+1, j)
			dfs(traveled, i-1, j)
			dfs(traveled, i, j+1)
			dfs(traveled, i, j-1)
		}
		return
	}

	var num int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}
			if traveled[i][j] {
				continue
			}
			dfs(traveled, i, j)
			num++
		}
	}
	return num

}
