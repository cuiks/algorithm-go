package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// root被监控数量
// 整棵树被监控数量
// 子树被监控数量

func minCameraCover(root *TreeNode) int {
	var dfs func(root *TreeNode) (a, b, c int)
	dfs = func(root *TreeNode) (a, b, c int) {
		if root == nil {
			return math.MaxInt32, 0, 0
		}
		la, lb, lc := dfs(root.Left)
		ra, rb, rc := dfs(root.Right)
		a = lc + rc + 1
		b = min(a, min(la+rb, lb+ra))
		c = min(a, lb+rb)
		return a, b, c
	}
	_, ans, _ := dfs(root)
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
