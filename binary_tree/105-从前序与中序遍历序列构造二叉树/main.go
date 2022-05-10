package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var findIndex func(target int) int
	findIndex = func(target int) int {
		for i, v := range inorder {
			if v == target {
				return i
			}
		}
		return -1
	}
	index := findIndex(preorder[0])
	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:index+1], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
	return root
}

func main() {

}
