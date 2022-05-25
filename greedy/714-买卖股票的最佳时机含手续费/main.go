package main

func maxProfit(prices []int, fee int) int {
	result := 0
	min := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > fee {
			result += prices[i] - min - fee
			min = prices[i] - fee
		}
	}
	return result
}

func main() {

}
