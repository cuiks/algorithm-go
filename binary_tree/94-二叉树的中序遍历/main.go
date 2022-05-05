package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
//func inorderTraversal(root *TreeNode) []int {
//	result := make([]int, 0)
//	var traversal func(root *TreeNode)
//	traversal = func(root *TreeNode) {
//		if root == nil {
//			return
//		}
//		traversal(root.Left)
//		result = append(result, root.Val)
//		traversal(root.Right)
//	}
//	traversal(root)
//	return result
//}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

func main() {

}
