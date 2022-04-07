package main

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := (start + end) / 2
		if target < nums[mid] {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	return -1
}

func main() {

}
