package main

func main() {

}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	row := len(image)
	if row <= 0 {
		return image
	}
	column := len(image[0])
	if column <= 0 {
		return image
	}

	if (sr < 0 || sr >= row) || (sc < 0 || sc >= column) {
		return image
	}

	oldColor := image[sr][sc]
	image = dfs(image, sr, sc, oldColor, newColor)
	//image = bfs(image, sr, sc, oldColor, newColor)

	return image
}

func dfs(image [][]int, sr int, sc int, oldColor int, newColor int) [][]int {
	row := len(image)
	column := len(image[0])
	if image[sr][sc] == oldColor {
		image[sr][sc] = newColor
		if sr-1 >= 0 {
			image = dfs(image, sr-1, sc, newColor, oldColor)
		}
		if sr+1 < row {
			image = dfs(image, sr+1, sc, newColor, oldColor)
		}
		if sc-1 >= 0 {
			image = dfs(image, sr, sc-1, newColor, oldColor)
		}
		if sc+1 < column {
			image = dfs(image, sr, sc+1, newColor, oldColor)
		}
	}
	return image
}

func bfs(image [][]int, sr int, sc int, oldColor int, newColor int) [][]int {
	row := len(image)
	column := len(image[0])

	if image[sr][sc] != oldColor {
		return image
	}
	image[sr][sc] = newColor

	if sr-1 >= 0 {
		image = bfs(image, sr-1, sc, oldColor, newColor)
	}
	if sr+1 < row {
		image = bfs(image, sr+1, sc, oldColor, newColor)
	}
	if sc-1 >= 0 {
		image = bfs(image, sr, sc-1, oldColor, newColor)
	}
	if sc+1 < column {
		image = bfs(image, sr, sc+1, oldColor, newColor)
	}

	return image

}
