package main

//https://leetcode.cn/problems/next-greater-element-i/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	dict := map[int]int{}
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			dict[num] = stack[len(stack)-1]
		} else {
			dict[num] = -1
		}
		stack = append(stack, num)
	}
	result := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		result[i] = dict[nums1[i]]
	}
	return result
}

func main() {

}
