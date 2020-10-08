package main

import "fmt"

func main() {
	nums := []int{2, 2, 2, 0, 1}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if min > nums[i] {
			return nums[i]
		}
		min = nums[i]
	}
	return nums[0]
}
