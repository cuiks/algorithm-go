package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxIndex := 0
	for i := 0; i < len(nums)-1; i++ {
		if maxIndex < i {
			return false
		}
		if v := i + nums[i]; v > maxIndex {
			maxIndex = v
		}
		if maxIndex >= len(nums)-1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))

}
