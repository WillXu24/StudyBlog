package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	nextPermutation(nums)

	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	length := len(nums)

	i, j, k := length-2, length-1, length-1

	for ; i >= 0 && nums[i] >= nums[j]; i, j = i-1, j-1 {
	}

	if i >= 0 {
		for ; k >= j && nums[i] >= nums[k]; k-- {
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	for m, n := j, length-1; m < n; m, n = m+1, n-1 {
		nums[m], nums[n] = nums[n], nums[m]
	}
}
