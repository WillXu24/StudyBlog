package main

import (
	"fmt"	
	"strings"
)


func main() {
	res:=lengthOfLongestSubstring("abcabcbb")
	fmt.Println(res)
}


func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
			continue
		}
		start = start + index + 1
		end = end + index + 1
	}
	return end - start
}