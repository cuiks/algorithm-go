package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		temp := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			queue = append(queue, node.Children...)
		}
		result = append(result, temp)
	}
	return result
}

func main() {

}
