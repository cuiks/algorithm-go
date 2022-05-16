package main

func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	backtrack(nums, 0, []int{}, &result)
	return result
}

func backtrack(nums []int, index int, list []int, result *[][]int) {
	if len(list) >= 2 {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
	}
	record := map[int]bool{}
	for i := index; i < len(nums); i++ {
		if (len(list) > 0 && nums[i] < list[len(list)-1]) || record[nums[i]] {
			continue
		}
		record[nums[i]] = true
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[:len(list)-1]
	}
}

func main() {

}
