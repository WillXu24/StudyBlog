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

// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
// 与主站 206 题相同
// 递归解法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList(head.Next)
	//next.Next = head
	head.Next.Next = head
	head.Next = nil
	return node
}
