package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	for _, v := range intervals {
		if len(result) == 0 || result[len(result)-1][1] < v[0] {
			result = append(result, v)
		} else {
			if v[1] > result[len(result)-1][1] {
				result[len(result)-1][1] = v[1]
			}
		}
	}
	return result
}

func main() {

}
