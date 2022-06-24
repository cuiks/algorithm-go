package main

// https://leetcode.cn/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dict := map[string]bool{}
	for _, v := range wordDict {
		dict[v] = true
	}
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {

}
