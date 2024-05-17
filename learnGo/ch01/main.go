package main

import "fmt"

// 定义在main函数外为全局变量
// 这里定义的全局变量没有用到不会报错
// var name = "tom"
// var ok bool //布尔类型默认值false
// 上面两行代码通常这样写：
var (
	name = "tom"
	ok   bool
)

//简洁变量定义不能作用于全局变量
//age:=1 //报错

func main() {
	//go是静态语言
	/*
		1、变量先定义后使用
		2、变量必须有类型
		3、类型定义后不能改变
		4、同一作用域 变量名不能冲突
		5、简洁变量定义不能作用于全局变量
		6、变量有默认值
	*/

	//变相定义方式：
	//var age int=1 //类型定义在后边
	//var age=1		//int可省略
	age := 100 //:=为赋值符号，可以省略var
	//go语言中局部变量定义后不使用会报错
	fmt.Println(age)

	//多变量定义
	//var user1, user2, user3 string //默认为空字符串
	var user1, user2, user3 = "tom1", 1, "tom3"
	fmt.Println(user1, user2, user3)

	/* 注意：

	 */
}
