package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ans := dp[0]
	for _, v := range dp {
		ans = max(ans, v)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
}
