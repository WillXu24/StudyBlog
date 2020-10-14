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
	// 伪头节点，见下面的关键点
	res := &ListNode{}
	temp := res
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			// 关键点：被赋值的是temp的下一节点而不是temp节点，因为temp变了，res就没用了！！！
			temp.Next, l1 = l1, l1.Next
		} else {
			temp.Next, l2 = l2, l2.Next
		}
		temp = temp.Next
	}
	if l1 == nil {
		temp.Next = l2
	} else {
		temp.Next = l1
	}
	return res.Next
}
