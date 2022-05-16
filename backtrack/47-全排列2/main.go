package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	backtrack(nums, map[int]bool{}, []int{}, &result)
	return result
}

func backtrack(nums []int, record map[int]bool, list []int, result *[][]int) {
	if len(nums) == len(list) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if record[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !record[i-1] {
			continue
		}
		record[i] = true
		list = append(list, nums[i])
		backtrack(nums, record, list, result)
		record[i] = false
		list = list[:len(list)-1]
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
