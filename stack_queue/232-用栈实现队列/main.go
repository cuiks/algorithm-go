package main

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	for len(this.stack2) != 0 {
		val := this.stack2[len(this.stack2)-1]
		this.stack1 = append(this.stack1, val)
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	for len(this.stack1) != 0 {
		val := this.stack1[len(this.stack1)-1]
		this.stack2 = append(this.stack2, val)
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
	key := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return key
}

func (this *MyQueue) Peek() int {
	for len(this.stack1) != 0 {
		val := this.stack1[len(this.stack1)-1]
		this.stack2 = append(this.stack2, val)
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
	return this.stack2[len(this.stack2)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.stack2) != 0 || len(this.stack1) != 0 {
		return false
	}
	return true
}

func main() {

}
