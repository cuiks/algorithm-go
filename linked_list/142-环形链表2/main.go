package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			fast = head
			slow = slow.Next
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}

func main() {

}
