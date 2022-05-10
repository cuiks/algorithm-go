package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := searchBST(root.Left, val); left != nil {
		return left
	}
	if right := searchBST(root.Right, val); right != nil {
		return right
	}
	return nil
}

func main() {

}
