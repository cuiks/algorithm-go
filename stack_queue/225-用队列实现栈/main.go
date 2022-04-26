package main

type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	for len(this.queue1) > 0 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	this.queue1 = append(this.queue1, x)
	for len(this.queue2) > 0 {
		this.queue1 = append(this.queue1, this.queue2[0])
		this.queue2 = this.queue2[1:]
	}
}

func (this *MyStack) Pop() int {
	val := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return val
}

func (this *MyStack) Top() int {
	return this.queue1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0 && len(this.queue2) == 0
}

func main() {

}
