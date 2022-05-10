package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	record := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		record = append(record, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	min := math.MaxInt64
	for i := 0; i < len(record)-1; i++ {
		if diff := abs(record[i], record[i+1]); diff < min {
			min = diff
		}
	}
	return min
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func main() {

}
