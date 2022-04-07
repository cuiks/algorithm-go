package main

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := (start + end) / 2
		if target < nums[mid] {
			end = mid
		} else if target > nums[mid] {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		result[0] = start
	} else if nums[end] == target {
		result[0] = end
	} else {
		result[0], result[1] = -1, -1
		return result
	}

	start, end = 0, len(nums)-1
	for start+1 < end {
		mid := (start + end) / 2
		if target < nums[mid] {
			end = mid
		} else if start > nums[mid] {
			start = mid
		} else {
			start = mid
		}
	}
	if nums[end] == target {
		result[1] = end
	} else if nums[start] == target {
		result[1] = start
	} else {
		result[0], result[1] = -1, -1
		return result
	}
	return result
}

func main() {

}
