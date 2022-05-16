package main

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	backtrack(nums, map[int]bool{}, []int{}, &result)
	return result
}

func backtrack(nums []int, record map[int]bool, list []int, result *[][]int) {
	if len(list) == len(nums) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if record[nums[i]] {
			continue
		}
		record[nums[i]] = true
		list = append(list, nums[i])
		backtrack(nums, record, list, result)
		record[nums[i]] = false
		list = list[:len(list)-1]
	}
}

func main() {

}
