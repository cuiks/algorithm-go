package main

import "fmt"

func package01(weight, value []int, packageWeight int) int {
	k := len(value)
	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, packageWeight+1)
	}
	for i := packageWeight; i >= weight[0]; i-- {
		dp[0][i] = dp[0][i-weight[0]] + value[0]
	}
	for i := 1; i < k; i++ {
		for j := weight[i]; j <= packageWeight; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
		}
	}
	return dp[k-1][packageWeight]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	fmt.Println(package01(weight, value, 4))
}
