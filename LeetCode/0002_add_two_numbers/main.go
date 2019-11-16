package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	res := addTwoNumbers(&l1, &l2)
	for ; res != nil; res = res.Next {
		defer fmt.Print(res.Val)
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return digAdd(l1, l2, 0)
}

func digAdd(digL1 *ListNode, digL2 *ListNode, carry int) *ListNode {
	// prepare part
	if digL1 == nil && digL2 == nil {
		if carry != 0 {
			return &ListNode{Val: carry}
		}
		return nil
	}
	if digL1 == nil {
		digL1 = &ListNode{Val: 0}
	}
	if digL2 == nil {
		digL2 = &ListNode{Val: 0}
	}
	// add part
	digSum := digL1.Val + digL2.Val + carry
	digRes := ListNode{Val: digSum % 10}
	carry = digSum / 10
	digRes.Next = digAdd(digL1.Next, digL2.Next, carry)
	return &digRes
}
