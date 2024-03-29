package main

/*
You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti.
You are also given an interval newInterval = [start, end] that represents the start and end of another interval.
Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).
Return intervals after the insertion.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
*/
func main() {

}

func insert(intervals [][]int, newInterval []int) [][]int {
	size := len(intervals)
	if size <= 0 {
		return [][]int{newInterval}
	}
	var rs [][]int

	mergeStart := newInterval[0]
	mergeEnd := newInterval[1]
	needAppend := true

	for i := 0; i < size; i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		if !needAppend {
			rs = append(rs, intervals[i])
			continue
		}

		if start < mergeStart && end < mergeStart {
			rs = append(rs, intervals[i])
		} else if mergeStart < start && mergeEnd < start {
			rs = append(rs, []int{mergeStart, mergeEnd})
			needAppend = false
			rs = append(rs, intervals[i])
		} else if mergeStart > end && mergeEnd > end {
			rs = append(rs, []int{mergeStart, mergeEnd})
			needAppend = false
			rs = append(rs, intervals[i])
		} else if start > mergeEnd && end > mergeEnd {
			rs = append(rs, []int{mergeStart, mergeEnd})
			needAppend = false
			rs = append(rs, intervals[i])
		} else {
			if mergeStart > start {
				mergeStart = start
			}
			if mergeEnd < end {
				mergeEnd = end
			}
			needAppend = true
		}
	}
	if needAppend {
		rs = append(rs, []int{mergeStart, mergeEnd})
	}

	return rs
}

func insert2(intervals [][]int, newInterval []int) [][]int {
	size := len(intervals)
	if size <= 0 {
		return [][]int{newInterval}
	}

	var rs [][]int

	for i, interval := range intervals {
		if newInterval[1] < interval[0] {
			rs = append(rs, newInterval)
			rs = append(rs, intervals[i:]...)
			return rs
		}

		if newInterval[0] > interval[1] {
			rs = append(rs, interval)
			continue
		}

		if newInterval[0] > interval[0] {
			newInterval[0] = interval[0]
		}
		if newInterval[1] < interval[1] {
			newInterval[1] = interval[1]
		}
	}
	return append(rs, newInterval)
}
