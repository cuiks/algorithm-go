package main

func maxProfit(prices []int) int {
	low := prices[0]
	res := 0
	for _, v := range prices {
		low = min(low, v)
		res = max(res, v-low)
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
