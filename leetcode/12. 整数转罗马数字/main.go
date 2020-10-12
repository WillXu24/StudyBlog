package main

import "fmt"

func main() {
	res := intToRoman(3)
	fmt.Println(res)
}

func intToRoman(num int) string {
	var res string
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	letters := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < 13; {
		if num >= values[i] {
			num -= values[i]
			res += letters[i]
		} else {
			i++
		}
	}
	return res
}
