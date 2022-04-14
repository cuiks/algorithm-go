package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	pre := slow
	for fast != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}
	if pre == slow {
		head = head.Next
	} else {
		pre.Next = slow.Next
	}
	return head
}

func main() {

}
