package main

import "fmt"

func main() {
	/*
		变量的作用域为定义的{}中
	*/

	{
		name := "hah"
		fmt.Println(name) //hah
	}
	//fmt.Println(name)//访问不到

	var age = 100
	if age == 100 {
		age := 1
		fmt.Println(age) //1
	}
	fmt.Println(age) //100
}
