package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var find func() int
	find = func() int {
		index := 0
		max := nums[0]
		for i, v := range nums {
			if v > max {
				max = v
				index = i
			}
		}
		return index
	}
	index := find()
	root := &TreeNode{
		Val:   nums[index],
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
	return root
}

func main() {

}
