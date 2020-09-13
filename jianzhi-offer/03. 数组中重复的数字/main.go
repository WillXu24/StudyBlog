package main

import "fmt"

func main() {
	res := findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3})
	fmt.Println(res)
}

// 关键：0~n-1所有数值都在
// 采用：原地置换
//
func findRepeatNumber(nums []int) int {
	for i := 0; i <= len(nums); i++ {
		for i != nums[i] {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return 0
}
