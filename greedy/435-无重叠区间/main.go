package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	result := 1
	max := intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= max {
			max = p[1]
			result++
		}
	}
	return len(intervals) - result
}

func main() {

}
