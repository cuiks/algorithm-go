package main

import "fmt"

func candy(ratings []int) int {
	left := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	result := 0
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			left[i] = max(left[i], left[i+1]+1)
		}
		result += left[i]
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}))

}
