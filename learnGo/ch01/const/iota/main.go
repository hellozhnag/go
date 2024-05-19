package main

import "fmt"

// 定义一个函数
func a() (int, bool) {
	return 0, false
}

func main() {
	// 匿名变量，定义后但是不想使用
	var _ int
	//比如这里的r1我们不关心，我们只关心返回ok值有没有成功
	//r1, ok := a()
	_, ok := a() //改成匿名变量接收
	if ok {
		//如果不使用 _ 匿名变量接收，那我们必须要使用下r1变量，否则报错
		//fmt.Println(r1)
	}

	// iota，特殊常量，可以认为是一个可以被编译器修改的常量

	/*
		1、后续会自动递增
		2、自增类型默认int
		3、每次出现const关键字时 iota归0
	*/
	const (
		err1 = iota
		err2
		err25
		err3
		err4 = "haha"
		err5
		err6 = iota
		err7 = 100
		err8 = iota
	)
	fmt.Println(err1, err2, err25, err3, err4, err5, err6, err7, err8)
	// 0 1 2 3 haha haha 6 100 8

	const (
		errnew = iota
	)
	fmt.Println(errnew) //0

	const ha = iota
	const ha1 = iota
	fmt.Println(ha, ha1) //0 0
}
