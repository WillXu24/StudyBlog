package main

import "fmt"

// MinStack 最小栈
type MinStack struct {
	Val []int
	Min []int
}

func main()  {
	obj := Constructor();
	obj.Push(-2);
	obj.Push(0);
	obj.Push(-3);

	fmt.Println(obj.GetMin())

	obj.Pop();
	fmt.Println(obj.Top());
	fmt.Println(obj.GetMin());
}

// Constructor 初始化
func Constructor() MinStack {
	return MinStack{
		Val: []int{},
		Min: []int{},
	}
}

// Push 入栈
func (this *MinStack) Push(x int)  {
	this.Val=append(this.Val,x)
	if len(this.Min)==0{
		this.Min=append(this.Min,x)
		return
	}
	min:=this.Min[len(this.Min)-1]
	if min>x{
		this.Min=append(this.Min,x)
	}else{
		this.Min=append(this.Min,min)
	}
}

// Pop 出栈
func (this *MinStack) Pop()  {
	length:=len(this.Min)-1
	this.Min=this.Min[:length]
	this.Val=this.Val[:length]
}

// Top 获取栈顶
func (this *MinStack) Top() int {
	return this.Val[len(this.Val)-1]
}

// GetMin 获取最小值
func (this *MinStack) GetMin() int {
	return this.Min[len(this.Val)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */