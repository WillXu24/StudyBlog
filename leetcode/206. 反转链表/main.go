package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	res := reverseList(head)
	fmt.Println(res.Val)
	fmt.Println(res.Next.Val)
}

// ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代解法：注意迭代条件与返回值！
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2 := head, head.Next
	for p2 != nil {
		temp := p2.Next
		p2.Next = p1
		p1, p2 = p2, temp
	}
	head.Next = nil
	//fmt.Println(p2,p1)
	//p2.Next=p1
	return p1
}
