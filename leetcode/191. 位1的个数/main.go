package main

import "fmt"

func main() {
	num := uint32(9)
	fmt.Println(hammingWeight(num))
}

func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		res += int(num - num>>1<<1)
		num = num >> 1
	}
	return res
}
