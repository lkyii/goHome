package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const fileName  = "./1.txt"

// 数字转二进制
func toBin(n int) (result string) {
	result = ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result += strconv.Itoa(lsb)
	}
	return
}

// 读文件 一行行读
func readFile(fileName string){

	file,err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	fileScan := bufio.NewScanner(file)

	for fileScan.Scan() {
		fmt.Println(fileScan.Text())
	}
}

func main() {
	fmt.Println(
		toBin(1),
		toBin(2),
		toBin(100),
		toBin(99),
		toBin(5999),
		)

	readFile(fileName)
}