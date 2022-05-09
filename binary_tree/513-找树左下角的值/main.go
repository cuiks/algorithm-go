package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func findBottomLeftValue(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	var result int
//	queue := make([]*TreeNode, 0)
//	queue = append(queue, root)
//	for len(queue) > 0 {
//		l := len(queue)
//		result = queue[0].Val
//		for i := 0; i < l; i++ {
//			node := queue[0]
//			queue = queue[1:]
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//		}
//	}
//	return result
//}

func findBottomLeftValue(root *TreeNode) int {
	var result int
	maxDeep := -1
	var dfs func(root *TreeNode, deep int)
	dfs = func(root *TreeNode, deep int) {
		if root == nil {
			return
		}
		if deep > maxDeep {
			maxDeep = deep
			result = root.Val
		}
		dfs(root.Left, deep+1)
		dfs(root.Right, deep+1)
	}
	dfs(root, 0)
	return result
}

func main() {

}
