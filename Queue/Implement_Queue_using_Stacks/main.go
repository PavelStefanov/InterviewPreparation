package main

import "fmt"

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	for len(this.stack1) > 0 {
		n := len(this.stack1) - 1 // Top element
		element := this.stack1[n]

		this.stack1 = this.stack1[:n] // Pop

		this.stack2 = append(this.stack2, element)
	}

	this.stack2 = append(this.stack2, x)

	for len(this.stack2) > 0 {
		n := len(this.stack2) - 1 // Top element
		element := this.stack2[n]

		this.stack2 = this.stack2[:n] // Pop

		this.stack1 = append(this.stack1, element)
	}
}

func (this *MyQueue) Pop() int {
	n := len(this.stack1) - 1
	element := this.stack1[n]

	this.stack1 = this.stack1[:n]

	return element
}

func (this *MyQueue) Peek() int {
	n := len(this.stack1) - 1
	return this.stack1[n]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)

	for !obj.Empty() {
		peek := obj.Peek()
		pop := obj.Pop()
		fmt.Println(peek)
		fmt.Println(pop)
	}
}
