package main

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := (start + end) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			end = mid
		} else {
			start = mid
		}
	}
	if target <= nums[start] {
		return start
	} else if target <= nums[end] {
		return end
	} else {
		return end + 1
	}
}

func main() {

}
