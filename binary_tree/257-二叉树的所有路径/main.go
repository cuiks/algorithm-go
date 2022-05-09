package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	if root == nil {
		return result
	}

	var dfs func(root *TreeNode, s string)
	dfs = func(root *TreeNode, s string) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			result = append(result, s+strconv.Itoa(root.Val))
			return
		}
		s += strconv.Itoa(root.Val) + "->"
		dfs(root.Left, s)
		dfs(root.Right, s)
	}
	dfs(root, "")
	return result
}

func main() {

}
