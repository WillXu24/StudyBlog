package main

import "fmt"

func main() {
	nums := []int{2, 2, 2, 0, 1}
	fmt.Println(minArray(nums))
}

func minArray(numbers []int) int {
	//min := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			return numbers[i]
		}
		//min = numbers[i]
	}
	return numbers[0]
}
