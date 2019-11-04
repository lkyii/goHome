package main

import "fmt"

func main(){
	m := map[string]string{
		"name":"k",
		"age":"1",
	}

	m1 := make(map[string]string)

	var m2 map[string]string

	for _,v := range m{
		fmt.Println(v)
	}

	fmt.Println(`get m name m["name"]：`,m["name"])

	if name,ok := m["name"];ok {
		fmt.Println(name)
	}
	
	fmt.Println(m,m1,m2)
}