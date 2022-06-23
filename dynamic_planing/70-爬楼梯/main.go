package main

//func climbStairs(n int) int {
//	dp := make([]int, n+1)
//	dp[0], dp[1] = 1, 1
//	for i := 2; i <= n; i++ {
//		dp[i] = dp[i-2] + dp[i-1]
//	}
//	return dp[n]
//}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= 2; j++ {
			if i >= j {
				dp[i] += dp[i-j]
			}
		}
	}
	return dp[n]
}

func main() {

}
