package main

import "fmt"

// https://leetcode.cn/problems/partition-equal-subset-sum/

//func canPartition(nums []int) bool {
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//	if sum%2 != 0 {
//		return false
//	}
//	target := sum / 2
//	dp := make([][]bool, len(nums))
//	for i := 0; i < len(dp); i++ {
//		dp[i] = make([]bool, target+1)
//		dp[i][0] = true
//	}
//
//	for i := 1; i < len(dp); i++ {
//		for j := 1; j <= target; j++ {
//			if j >= nums[i] {
//				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
//			} else {
//				dp[i][j] = dp[i-1][j]
//			}
//		}
//	}
//	return dp[len(nums)-1][target]
//}

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, v := range nums {
		for i := target; i >= v; i-- {
			dp[i] = max(dp[i], dp[i-v]+v)
		}
	}
	return dp[target] == target
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
