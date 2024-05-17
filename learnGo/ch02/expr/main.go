package main

import "fmt"

func main() {
	var a, b = 1, 2
	fmt.Println(a + b)
	a++
	fmt.Println(a)
	var astr, bstr = "tom,", "hello"
	fmt.Println(astr + bstr)

	var A, B = 60, 13
	fmt.Println(A ^ B) //49
}
