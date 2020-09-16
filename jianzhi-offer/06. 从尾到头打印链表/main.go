package main

import "fmt"

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	res := reversePrint(&head)
	fmt.Println(res)
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：4.7 MB, 在所有 Go 提交中击败了13.84%的用户
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	return reverse(head, []int{})
}

func reverse(l *ListNode, nums []int) []int {
	if l == nil {
		return nums
	}
	res := reverse(l.Next, nums)
	nums = append(res, l.Val)
	return nums
}
