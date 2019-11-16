package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	maxRes := 0
	area := 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] > height[j] {
			area = (j - i) * height[j]
			j--
		} else {
			area = (j - i) * height[i]
			i++
		}
		if area > maxRes {
			maxRes = area
		}
	}
	return maxRes
}
