package main

import "sort"

/*
*
https://leetcode.com/problems/k-closest-points-to-origin/
Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).
The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).
You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).
*/
func main() {

}

func kClosest(points [][]int, k int) [][]int {
	if k == len(points) {
		return points
	}

	for i, point := range points {
		points[i] = append(points[i], distance(point[0], point[1]))
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][2] < points[j][2]
	})

	rs := [][]int{}
	for i := 0; i < k; i++ {
		rs = append(rs, points[i][:2])
	}
	return rs
}

func distance(x, y int) int {
	return x*x + y*y
}
