package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	result := 0
	for i := 1; i < len(prices); i++ {
		if v := prices[i] - prices[i-1]; v > 0 {
			result += v
		}
	}
	return result
}

func main() {

}
