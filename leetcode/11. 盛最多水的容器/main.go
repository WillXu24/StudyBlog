package main

import "fmt"

func main() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(h)
	fmt.Println(res)
}

func maxArea(height []int) int {
	var res, temp int
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			temp = height[i] * (j - i)
			i++
		} else {
			temp = height[j] * (j - i)
			j--
		}
		if temp > res {
			res = temp
		}
	}
	return res
}
