package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+nums[i] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
