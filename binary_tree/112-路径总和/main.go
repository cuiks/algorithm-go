package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	result := false
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		if result == true || (root.Left == nil && root.Right == nil && sum == targetSum) {
			result = true
			return
		}
		dfs(root.Left, sum)
		dfs(root.Right, sum)
	}
	dfs(root, 0)
	return result
}

func main() {

}
