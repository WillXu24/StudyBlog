package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
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
