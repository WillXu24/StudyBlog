package main

import "fmt"

func main() {
	res := convert("LEETCODEISHIRING", 3)
	fmt.Println(res)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var res []byte
	numSection := 2*numRows - 2
	length := len(s)
	// first row
	for i := 0; i < length; i += numSection {
		res = append(res, s[i])
	}

	// middle rows
	for i := 1; i < numRows-1; i++ {
		for j := i; j < length; j += numSection {
			res = append(res, s[j])
			if j+2*(numRows-1-i) < length {
				res = append(res, s[j+2*(numRows-1-i)])
			}
		}
	}

	// last row
	for i := numRows - 1; i < length; i += numSection {
		res = append(res, s[i])
	}

	return string(res)
}
