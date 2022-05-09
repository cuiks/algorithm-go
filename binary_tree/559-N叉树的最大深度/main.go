package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	for _, node := range root.Children {
		if cur := maxDepth(node); cur > max {
			max = cur
		}
	}
	return max + 1
}

func main() {

}
