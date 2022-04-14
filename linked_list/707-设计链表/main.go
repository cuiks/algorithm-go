package main

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Dummy *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Dummy: &Node{},
	}
}

func (this *MyLinkedList) Get(index int) int {
	dummy := this.Dummy
	head := dummy.Next
	i := 0
	for head != nil {
		if i == index {
			return head.Val
		}
		i++
		head = head.Next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.Dummy
	head := dummy.Next
	node := &Node{Val: val, Next: head}
	dummy.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}
	dummy := this.Dummy
	for dummy.Next != nil {
		dummy = dummy.Next
	}
	dummy.Next = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	dummy := this.Dummy
	for dummy != nil && index > 0 {
		dummy = dummy.Next
		index--
	}
	if dummy == nil {
		return
	}
	node := &Node{Val: val, Next: dummy.Next}
	dummy.Next = node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	dummy := this.Dummy
	for i := 0; i < index; i++ {
		if dummy.Next == nil {
			return
		}
		dummy = dummy.Next
	}
	if dummy.Next != nil {
		dummy.Next = dummy.Next.Next
	}
}

/*
["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
[[],[1],[3],[1,2],[1],[1],[1]]
*/

func main() {
	l := Constructor()
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 2)
	l.Get(1)
	l.DeleteAtIndex(1)
	l.Get(1)
	//l.AddAtHead(2)
	//l.DeleteAtIndex(1)
	//l.AddAtHead(2)
	//l.AddAtHead(7)
	//l.AddAtHead(3)
	//l.AddAtHead(2)
	//l.AddAtHead(5)
	//l.AddAtTail(5)
	//l.Get(5)
	//l.DeleteAtIndex(6)
	//l.DeleteAtIndex(4)
	//l.AddAtIndex(1, 0)
	//l.Get(0)
}
