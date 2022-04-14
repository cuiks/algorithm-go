package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		left := cur.Next
		right := cur.Next.Next
		cur.Next = right
		left.Next = right.Next
		right.Next = left
		cur = cur.Next.Next
	}
	return dummy.Next
}

func main() {

}
