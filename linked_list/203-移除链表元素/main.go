package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	head = dummy
	for head != nil && head.Next != nil {
		for head.Next != nil && head.Next.Val == val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return dummy.Next
}

func main() {

}
