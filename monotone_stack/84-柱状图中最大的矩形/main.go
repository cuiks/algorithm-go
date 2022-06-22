package main

import "fmt"

// https://leetcode.cn/problems/largest-rectangle-in-histogram/

func largestRectangleArea(heights []int) int {
	res := 0
	cur := 0
	stack := make([]int, 0)
	for i := 0; i <= len(heights); i++ {
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}
		for len(stack) > 0 && cur <= heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			}
			res = max(res, h*w)
		}
		stack = append(stack, i)
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
