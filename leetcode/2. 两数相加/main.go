/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import "fmt"

// ListNode 链表
type ListNode struct {
	     Val int
	     Next *ListNode
}

func main() {
l1:=ListNode{Val: 2,Next:&ListNode{Val:4,Next:&ListNode{Val:3,Next:nil}}}
l2:=ListNode{Val: 5,Next:&ListNode{Val:6,Next:&ListNode{Val:4,Next:nil}}}

res:=addTwoNumbers(&l1,&l2)
for res.Next!=nil{
	fmt.Println(res.Val)
	res=res.Next
}
fmt.Println(res.Val)
}


// 解法
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return digAdd(l1, l2, 0)
}

func digAdd(digL1 ,digL2 *ListNode, carry int) *ListNode {
    // 设定终止条件，注意最后的进位
    if digL1 == nil && digL2 == nil {
		if carry != 0 {
			return &ListNode{Val: carry}
		}
		return nil
	}
    // 防止数位不同
    if digL1 == nil {
		digL1 = &ListNode{Val: 0}
	}
	if digL2 == nil {
		digL2 = &ListNode{Val: 0}
	}
    // 迭代部分
    digSum := digL1.Val + digL2.Val + carry
	digRes := ListNode{Val: digSum % 10}
	carry = digSum / 10
	digRes.Next = digAdd(digL1.Next, digL2.Next, carry)
	return &digRes
}