package main

func findLengthOfLCIS(nums []int) int {
	ans := 1
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			index = i
		}
		if diff := i - index + 1; diff > ans {
			ans = diff
		}
	}
	return ans
}

func main() {

}
