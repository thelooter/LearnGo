package main

import "fmt"

func main() {

	var a = 27
	var b = &a
	a++

	fmt.Println("a = ", a)
	fmt.Println("b = ", *b)

	str := new(myStruct)
	str.a = 10
	str.b = 20

	newStr := *str
	newStr.a = 100
	(*str).a = 200

	fmt.Println("str = ", str.a)
	fmt.Println("newStr = ", (newStr).a)
}

type myStruct struct {
	a int
	b int
}
