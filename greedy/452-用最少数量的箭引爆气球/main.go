package main

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	max := points[0][1]
	result := 1
	for _, p := range points {
		if p[0] > max {
			result++
			max = p[1]
		}
	}
	return result
}

func main() {

}
