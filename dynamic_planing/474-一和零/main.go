package main

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		zero, one := 0, 0
		for _, v := range strs[i] {
			if v == '0' {
				zero++
			} else {
				one++
			}
		}
		for j := m; j >= zero; j-- {
			for k := n; k >= one; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zero][k-one]+1)
			}
		}
	}
	return dp[m][n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
