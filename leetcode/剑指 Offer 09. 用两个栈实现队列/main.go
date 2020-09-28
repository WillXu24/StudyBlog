package main

import (
	"fmt"
	"sync"
)

// CQueue 队列
type CQueue struct {
	stackA Stack
	stackB Stack
}

func main() {
	a := Constructor()
	a.AppendTail(3)
	res := a.DeleteHead()
	fmt.Println(res)
	res = a.DeleteHead()
	fmt.Println(res)
}

// 执行用时：248 ms, 在所有 Go 提交中击败了64.60%的用户
// 内存消耗：8.6 MB, 在所有 Go 提交中击败了16.63%的用户

// Constructor 建造
func Constructor() CQueue {
	return CQueue{
		stackA: *NewStack(),
		stackB: *NewStack(),
	}
}

// AppendTail 入列
func (this *CQueue) AppendTail(value int) {
	this.stackA.Push(value)
}

// DeleteHead 出列
func (this *CQueue) DeleteHead() int {
	lengthA, lengthB := len(this.stackA.val), len(this.stackB.val)
	if lengthA == 0 && lengthB == 0 {
		return -1
	}
	if lengthB == 0 {
		val := this.stackA.Pop()
		for val != -1 {
			this.stackB.Push(val)
			val = this.stackA.Pop()
		}
		return this.stackB.Pop()
	} else {
		return this.stackB.Pop()
	}
}

// Stack 栈
type Stack struct {
	val  []int
	lock sync.RWMutex
}

// NewStack 初始化栈
func NewStack() *Stack {
	return &Stack{
		val: []int{},
		//lock: sync.RWMutex{},
	}
}

// Print prints all the elements
func (s *Stack) Print() {
	fmt.Println(s.val)
}

// Push adds an Item to the top of the stack
func (s *Stack) Push(t int) {
	s.lock.Lock()
	s.lock.Unlock()
	s.val = append(s.val, t)
}

// Pop removes an Item from the top of the stack
func (s *Stack) Pop() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.val) == 0 {
		return -1
	}
	val := s.val[len(s.val)-1]
	s.val = s.val[0 : len(s.val)-1]
	return val
}
