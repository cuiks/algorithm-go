package main

func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] = min(cost[i-1], cost[i-2]) + cost[i]
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {

}
