package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i-1] + nums[i]
		}
		if nums[i] > result {
			result = nums[i]
		}
	}
	return result
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
