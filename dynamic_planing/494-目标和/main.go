package main

// https://leetcode.cn/problems/target-sum/

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum-target < 0 || (sum-target)%2 != 0 {
		return 0
	}
	bag := (sum - target) / 2
	dp := make([]int, bag+1)
	dp[0] = 1
	for _, v := range nums {
		for i := bag; i >= v; i-- {
			dp[i] = dp[i] + dp[i-v]
		}
	}
	return dp[bag]
}

func main() {

}
