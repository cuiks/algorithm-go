package main

import "fmt"

// https://leetcode.cn/problems/partition-equal-subset-sum/

//func canPartition(nums []int) bool {
//	sum := 0
//	for _, k := range nums {
//		sum += k
//	}
//	if sum%2 != 0 {
//		return false
//	}
//
//	target := sum / 2
//	dp := make([][]bool, len(nums))
//	for i, _ := range dp {
//		dp[i] = make([]bool, target+1)
//		dp[i][0] = true
//	}
//	for i := 1; i < len(nums); i++ {
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
	for _, num := range nums {
		for j := target; j >= num; j-- {
			if dp[j] < dp[j-num]+num {
				dp[j] = dp[j-num] + num
			}
		}
	}
	return dp[target] == target
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
