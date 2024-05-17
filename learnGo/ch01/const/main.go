package main

import "fmt"

func main() {
	//显式定义常量
	const PI float64 = 3.1415926
	//隐式定义
	const PI2 = 3.1415926

	//命名：大写，单词间使用_连接

	//可以一组常量定义
	const (
		UNKNOW = 1
		FEMALE = 2
		MALE   = 3
	)

	const (
		x int = 16
		y
		s = "abc"
		z
	)
	fmt.Println(x, y, s, z) //16 16 abc abc
	/* 默认规则：
	1、常量类型定义只支持 bool、数值（整数、浮点数、复数）、字符串
	2、不曾使用的常量 编译不报错
	3、定义时若没有设置类型和值，沿用前边
	*/
}
