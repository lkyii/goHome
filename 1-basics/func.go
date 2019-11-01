package main

import (
	"fmt"
)

// 函数返回的两个值
func countNum(a, b int, op string) (result int, err error) {
	switch {
	case op == "+":
		return a+b,nil
	case op == "-":
		return a+b,nil
	case op == "/":
		return a+b,nil
	case op == "*":
		return a+b,nil
	default:
		return 0,fmt.Errorf("undefind %s", op)
	}
}

func main()  {
	if result,err := countNum (1,2,"++"); err == nil{
		fmt.Println(result)
	}else{
		panic(err)
	}
}
