package main

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[j], j)*(i-j))
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
