package main

// https://leetcode.cn/problems/last-stone-weight-ii/

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}

	target := sum / 2
	dp := make([]int, target+1)
	for _, v := range stones {
		for i := target; i >= v; i-- {
			dp[i] = max(dp[i], dp[i-v]+v)
		}
	}
	return sum - 2*dp[target]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
