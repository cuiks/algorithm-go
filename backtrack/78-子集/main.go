package main

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	backtrack(nums, 0, []int{}, &result)
	return result
}

func backtrack(nums []int, index int, list []int, result *[][]int) {
	tmp := make([]int, len(list))
	copy(tmp, list)
	*result = append(*result, tmp)
	if index >= len(nums) {
		return
	}
	for i := index; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[:len(list)-1]
	}
}

func main() {
}
