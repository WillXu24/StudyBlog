package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	var res string
	if num < 1 || num > 39999 {
		return res
	}
	intValue := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanValue := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(intValue); {
		if num >= intValue[i] {
			num -= intValue[i]
			res += romanValue[i]
		} else {
			i++
		}
	}
	return res
}
