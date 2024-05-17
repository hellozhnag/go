package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 数值类型相互转换
	var a int = 12
	var b = uint(a)

	var f float32 = 3.14
	var c = int32(f)
	fmt.Println(b, c) //12,3

	var f64 = float64(a)
	fmt.Println(f64) //12

	// go语言类型要求严格
	type IT int
	//var d IT = a//报错
	var d IT = IT(a)
	fmt.Println(d)

	//字符、数字 相互转换
	var istr = "12"
	myint, err := strconv.Atoi(istr) //字符串转数字可能出错
	if err != nil {
		fmt.Println("convert err!")
	}
	fmt.Println(myint) //12

	mstr := strconv.Itoa(myint) //数字转字符串不会出错
	fmt.Println(mstr)           //12

	//一、字符串类型转其他基本类型
	// 1、字符串转换为float
	//这个地方如果写32会按照float32转换，但是函数仍然返回float64
	myf, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		fmt.Println("convert err!")
	}
	fmt.Println(myf) //3.14

	//2、字符串转int
	// 8进制，64位
	parseInt, err := strconv.ParseInt("12", 8, 64)
	if err != nil {
		fmt.Println("convert err!")
	}
	fmt.Println(parseInt) //10,12=8^0*2+8^1*1=10

	// 3、字符串转bool
	parseBool1, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("convert err!")
	}
	fmt.Println(parseBool1) //true
	parseBool2, err := strconv.ParseBool("hello")
	if err != nil {
		fmt.Println("convert err!") //convert err!
	}
	fmt.Println(parseBool2) //false,转换失败的话，为默认值false
	parseBool3, err := strconv.ParseBool("1")
	if err != nil {
		fmt.Println("convert err!")
	}
	fmt.Println(parseBool3) //true,"1"->true
	parseBool4, err := strconv.ParseBool("0")
	if err != nil {
		fmt.Println("convert err!")
	}
	fmt.Println(parseBool4) //false，"0"->false
	parseBool5, err := strconv.ParseBool("12")
	if err != nil {
		fmt.Println("convert err!") //convert err!
	}
	fmt.Println(parseBool5) //false,转换失败

	//二、基本类型转字符串
	formatBool := strconv.FormatBool(true)
	fmt.Println(formatBool) //true

	float := strconv.FormatFloat(3.14, 'E', -1, 64)
	fmt.Println(float) //3.14E+00

	//转换为16进制
	formatInt := strconv.FormatInt(42, 16)
	fmt.Println(formatInt) //2a,16^0*10+16^1*2=42

}
