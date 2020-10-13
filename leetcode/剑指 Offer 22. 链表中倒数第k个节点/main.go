package main

import "fmt"

func main() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	res := getKthFromEnd(&head, 2)
	fmt.Println(res)
}

// ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fastP := head
	for i := 1; i <= k; i++ {
		fastP = fastP.Next
	}
	//fmt.Println(fastP.Val)
	for fastP != nil {
		fastP = fastP.Next
		head = head.Next
	}
	return head
}
