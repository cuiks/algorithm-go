package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
//func preorderTraversal(root *TreeNode) []int {
//	result := make([]int, 0)
//	var traversal func(node *TreeNode)
//	traversal = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		result = append(result, node.Val)
//		traversal(node.Left)
//		traversal(node.Right)
//	}
//  traversal(root)
//	return result
//}

// 迭代
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

func main() {

}
