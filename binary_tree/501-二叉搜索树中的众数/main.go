package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	record := map[int]int{}
	max := 1

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		record[root.Val]++
		if record[root.Val] == max {
			result = append(result, root.Val)
		} else if record[root.Val] > max {
			max = record[root.Val]
			result = []int{root.Val}
		}
		dfs(root.Right)
	}
	dfs(root)
	return result
}

func main() {

}
