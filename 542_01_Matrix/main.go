package main

/*
https://leetcode.com/problems/01-matrix/

Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
The distance between two adjacent cells is 1.

Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
Output: [[0,0,0],[0,1,0],[0,0,0]]

Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
Output: [[0,0,0],[0,1,0],[1,2,1]]
*/
func main() {

}

func updateMatrix(mat [][]int) [][]int {

	var queue [][]int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				point := []int{i, j}
				queue = append(queue, point)
			} else {
				mat[i][j] = -1
			}
		}
	}

	dir := [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}
	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]

		//BFS
		for _, ele := range dir {
			x := point[0] + ele[0]
			y := point[1] + ele[1]
			if x >= 0 && x < len(mat) && y >= 0 && y < len(mat[0]) && mat[x][y] == -1 {
				mat[x][y] = mat[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}
	return mat
}
