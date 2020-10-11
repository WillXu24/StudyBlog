package main

import "fmt"

func main() {
	fmt.Println(printNumbers(1))
}

func printNumbers(n int) []int {
	if n == 0 {
		return []int{}
	}

	cap := 1
	for i := 0; i < n; i++ {
		cap *= 10
	}
	res := make([]int, cap-1)
	for i := range res {
		res[i] = i + 1
	}
	// res := []int{}
	// for i := 1; i < cap; i++ {
	// 	res = append(res, i)
	// }
	//fmt.Println(res)
	return res
}
