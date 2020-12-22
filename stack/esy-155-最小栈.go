package main

import "fmt"

// 关键是对于stack[1...N], minStack[top]维持的是1...N的最小值
type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {

	this.stack = append(this.stack, x)

	minStackLen := len(this.minStack)

	if minStackLen == 0 {
		this.minStack = append(this.minStack, x)
		return
	}

	if this.minStack[minStackLen-1] >= x {
		this.minStack = append(this.minStack, x)
		return
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if val == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}

}

func (this *MinStack) Top() int {

	if len(this.stack) == 0 {
		return 0
	}

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {

	if len(this.stack) == 0 {
		return 0
	}

	return this.minStack[len(this.minStack)-1]

}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()
	fmt.Println(param_3, param_4)

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
