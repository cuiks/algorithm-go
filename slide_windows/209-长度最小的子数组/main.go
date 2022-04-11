package main

func minSubArrayLen(target int, nums []int) int {
	start, end := 0, 0
	sum := 0
	result := len(nums) + 1
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			if end-start+1 < result {
				result = end - start + 1
			}
			sum -= nums[start]
			start++
		}
		end++
	}
	if result < len(nums)+1 {
		return result
	}
	return 0
}

func main() {
}
