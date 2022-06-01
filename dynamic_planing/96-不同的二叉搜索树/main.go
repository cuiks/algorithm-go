package main

// 1 [2,3,4,5,6,7]
// [1] 2 [3,4,5,6,7]
// [1,2] 3 [4,5,6,7]

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func main() {

}
