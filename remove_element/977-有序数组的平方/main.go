package main

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	index := len(nums) - 1
	start, end := 0, len(nums)-1
	for start <= end {
		left := nums[start] * nums[start]
		right := nums[end] * nums[end]
		if left > right {
			result[index] = left
			start++
		} else {
			result[index] = right
			end--
		}
		index--
	}
	return result
}

func main() {

}
