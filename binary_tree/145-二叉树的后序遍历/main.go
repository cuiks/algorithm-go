package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func postorderTraversal(root *TreeNode) []int {
//	result := make([]int, 0)
//	var traversal func(root *TreeNode)
//	traversal = func(root *TreeNode) {
//		if root == nil {
//			return
//		}
//		traversal(root.Left)
//		traversal(root.Right)
//		result = append(result, root.Val)
//	}
//	traversal(root)
//	return result
//}

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	var last *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == last {
			result = append(result, node.Val)
			stack = stack[:len(stack)-1]
			last = node
		} else {
			root = node.Right
		}
	}
	return result
}

func main() {

}
