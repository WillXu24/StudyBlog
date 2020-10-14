package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}

	res := mergeTwoLists(l1, l2)
	fmt.Println(res.Val, res.Next.Val)
	// res = mergeTwoLists1(l1, l2)
	// fmt.Println(res, res.Next)
}

// ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l2.Val < l1.Val {
		l1, l2 = l2, l1
	}
	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}
