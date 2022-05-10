package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	Valid bool
	Max   *TreeNode
	Min   *TreeNode
}

func helper(root *TreeNode) BST {
	if root == nil {
		return BST{Valid: true}
	}
	left := helper(root.Left)
	right := helper(root.Right)
	res := BST{
		Valid: true,
		Max:   root,
		Min:   root,
	}
	if !left.Valid || !right.Valid {
		res.Valid = false
		return res
	}
	if left.Max != nil && left.Max.Val >= root.Val {
		res.Valid = false
		return res
	}
	if right.Min != nil && right.Min.Val <= root.Val {
		res.Valid = false
		return res
	}
	if left.Min != nil {
		res.Min = left.Min
	}
	if right.Max != nil {
		res.Max = right.Max
	}
	return res
}

func isValidBST(root *TreeNode) bool {
	return helper(root).Valid
}

func main() {

}
