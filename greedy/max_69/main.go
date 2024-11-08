package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximum69Number(9669))
}

func maximum69Number(num int) int {

	str := strconv.Itoa(num)
	for i := 0; i < len(str); i++ {
		if str[i] == '6' {
			return num + 3*pow(10, len(str)-i-1)
		}
	}
	return num
}

func pow(num, n int) int {
	if n == 0 {
		return 1
	}
	var result = num
	for i := 1; i < n; i++ {
		result *= num
	}
	return result
}
