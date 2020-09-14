package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := majorityElement(nums)
	fmt.Println(res)
}

func majorityElement(nums []int) int {
	a := map[int]int{}
	b, c := 0, 0
	for _, v := range nums {
		a[v]++
		if a[v] > b {
			b, c = a[v], v
		}
	}
	return c
}