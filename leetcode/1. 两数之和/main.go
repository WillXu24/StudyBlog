package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum1(nums, target))
	fmt.Println(twoSum2(nums, target))
}

// 暴力
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 哈希
func twoSum2(nums []int, target int) []int {
	numMap := map[int]int{}
	for i, v := range nums {
		if k, ok := numMap[target-v]; ok {
			return []int{k, i}
		}
		numMap[v] = i
	}
	return nil
}
