package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	res := exchange(nums)
	fmt.Println(res)
}

func exchange(nums []int) []int {
	var res, even []int
	for i := range nums {
		if nums[i]%2 == 0 {
			even = append(even, nums[i])
		} else {
			res = append(res, nums[i])
		}
	}
	res = append(res, even[:]...)
	return res
}
