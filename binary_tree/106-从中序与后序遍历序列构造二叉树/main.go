package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
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
	index := findIndex(postorder[len(postorder)-1])
	root := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(inorder[:index], postorder[:index]),
		Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1]),
	}
	return root
}

func main() {

}

// inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
