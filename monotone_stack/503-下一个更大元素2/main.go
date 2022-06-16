package main

func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = -1
	}
	stack := make([]int, 0)
	for i := 0; i < len(nums)*2; i++ {
		for len(stack) > 0 && nums[i%len(nums)] > nums[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[index] = nums[i%len(nums)]
		}
		stack = append(stack, i%len(nums))
	}
	return result
}

func main() {
	nextGreaterElements([]int{1, 2})
}
