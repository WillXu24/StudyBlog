package main

import "fmt"

func main() {
	a := Constructor()
	a.Push(1)
	a.Push(0)
	res := a.Min()
	fmt.Println(res)
	a.Pop()
	res = a.Min()
	fmt.Println(res)
}

// 同155题

// MinStack 最小栈
type MinStack struct {
	Val    []int
	MinVal []int
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{}
}

// Push 入栈
func (this *MinStack) Push(x int) {
	this.Val = append(this.Val, x)
	length := len(this.MinVal)
	if length == 0 {
		this.MinVal = append(this.MinVal, x)

	} else {
		min := this.MinVal[length-1]
		if x < min {
			this.MinVal = append(this.MinVal, x)
		} else {
			this.MinVal = append(this.MinVal, min)
		}
	}
}

func (this *MinStack) Pop() {
	this.Val = this.Val[:len(this.Val)-1]
	this.MinVal = this.MinVal[:len(this.MinVal)-1]
}

func (this *MinStack) Top() int {
	return this.Val[len(this.Val)-1]
}

func (this *MinStack) Min() int {
	return this.MinVal[len(this.MinVal)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
