package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	res := dfs(root)
	return max(res[0], res[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l := dfs(node.Left)
	r := dfs(node.Right)
	res := make([]int, 2)
	// 偷node房子
	res[0] = node.Val + l[1] + r[1]
	// 不偷node房子
	res[1] = max(l[0], l[1]) + max(r[0], r[1])
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
