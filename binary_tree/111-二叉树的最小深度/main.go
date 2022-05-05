package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	min := math.MaxInt64
	if root.Left != nil {
		left := minDepth(root.Left)
		if left < min {
			min = left
		}
	}
	if root.Right != nil {
		right := minDepth(root.Right)
		if right < min {
			min = right
		}
	}
	return min + 1
}

func main() {

}
