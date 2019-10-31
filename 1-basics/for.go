package main

import (
	"fmt"
	"strconv"
)

// 数字转二进制
func toBin(n int) (result string) {
	result = ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result += strconv.Itoa(lsb)
	}
	return
}

func main() {
	fmt.Println(
		toBin(1),
		toBin(2),
		toBin(100),
		toBin(99),
		toBin(5999),
		)
}