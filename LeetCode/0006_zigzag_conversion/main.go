package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("ASDASDASD", 2))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	length := 2*numRows - 2
	var res []byte
	// first row
	for i := 0; i < len(s); i += length {
		res = append(res, s[i])
	}
	//fmt.Println(string(res))
	// middle row
	for row := 1; row < numRows-1; row++ {
		// each row get 2 value:header & end
		for i := row; i < len(s); i += length {
			// get header
			res = append(res, s[i])
			// get end
			if i+2*(numRows-1-row) < len(s) {
				res = append(res, s[i+2*(numRows-1-row)])
			}
		}
	}
	//fmt.Println(string(res))
	// last row
	for i := numRows - 1; i < len(s); i += length {
		res = append(res, s[i])
	}
	//fmt.Println(string(res))
	//fmt.Println(len(res))
	return string(res)
}

func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	length := 2*numRows - 2
	var res = ""
	// first row
	for i := 0; i < len(s); i += length {
		res += string(s[i])
	}
	fmt.Println(res)
	// middle row
	for row := 1; row < numRows-1; row++ {
		// each row get 2 value:header & end
		for i := row; i < len(s); i += length {
			// get header
			res += string(s[i])
			// get end
			if i+2*(numRows-1-row) < len(s) {
				res += string(s[i+2*(numRows-1-row)])
			}
		}
	}
	fmt.Println(string(res))
	// last row
	for i := numRows - 1; i < len(s); i += length {
		res += string(s[i])
	}
	return res
}
