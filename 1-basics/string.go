package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main(){
	s := "Yes刘宽宇"
	fmt.Println(len(s))

	//for _,b := range []byte(s){
	//	fmt.Println(b)
	//}

	fmt.Println(utf8.RuneCountInString(s))

	byte := []byte(s)

	for len(byte) > 0 {
		ch,size := utf8.DecodeRune(byte)
		byte = byte[size:]
		fmt.Printf("%c",ch)
	}

	rune := []rune(s)

	for i, ch :=range rune  {
		fmt.Printf(" (%d,%c) ", i, ch)
	}

	str1 := "a b ,cd ac a"
	arr1 :=strings.Fields(str1)
	arr2 :=strings.Split(str1,",")

	fmt.Println("\n\nFields")
	fmt.Println("arr1 count",len(arr1),arr1)
	for k,v := range arr1 {
		fmt.Printf("[%d:%s]",k,v)
	}

	fmt.Println("\n\nSplit")
	fmt.Println("arr2 count",len(arr2),arr2)
	for k,v := range arr2 {
		fmt.Printf("[%d:%s]",k,v)
	}

	fmt.Println("\n\nJoin")
	string1 := strings.Join(arr1,"--")
	string2 := strings.Join(arr2,"--")
	fmt.Println(string1)
	fmt.Println(string2)

}