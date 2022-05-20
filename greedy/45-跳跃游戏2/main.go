package main

func jump(nums []int) int {
	maxPos := 0
	end := 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, nums[i]+i)
		if i == end {
			end = maxPos
			step++
		}
	}
	return step
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
