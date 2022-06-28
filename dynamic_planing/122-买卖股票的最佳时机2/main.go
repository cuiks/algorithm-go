package main

func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}

func main() {

}
