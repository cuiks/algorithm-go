package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		temp := make([]*Node, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		for i := 0; i < len(temp)-1; i++ {
			temp[i].Next = temp[i+1]
		}
	}
	return root
}

func main() {

}
