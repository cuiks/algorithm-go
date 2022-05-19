package main

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return 1
	}

	var pre, cur int
	result := 1
	for i := 0; i < len(nums)-1; i++ {
		cur = nums[i] - nums[i+1]
		if (pre <= 0 && cur > 0) || (pre >= 0 && cur < 0) {
			result++
			pre = cur
		}
	}
	return result
}

func main() {

}
