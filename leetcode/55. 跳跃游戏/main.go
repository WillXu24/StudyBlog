package main

import "fmt"

func main() {
	nums:=[]int{2,3,1,1,4}
	nums=[]int{1}
	res:=canJump(nums)
	fmt.Println(res)
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	// 最大可到达位置
	max := 0
	for i, v := range nums {
		if max >= i && i+v > max {
			max = i + v
		}
	}
	return max >= len(nums)-1
}
