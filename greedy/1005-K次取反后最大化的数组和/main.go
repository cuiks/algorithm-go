package main

import (
	"fmt"
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))

}
