package main

func moveZeroes(nums []int) {
	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[start] = nums[start], nums[i]
			start++
		}
	}
}
func main() {

}
