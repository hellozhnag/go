package main

import "fmt"

func main() {
	var a int8
	var b int16
	var c int32
	var d int64
	fmt.Println(a, b, c, d)
	var e int //动态类型,根据我们操作系统位数来定义
	fmt.Println(e)

	//a=b //报错
	//b=a//报错
	a = int8(b)
	b = int16(a)

	var f1 float32 //最大是3.4e38
	var f2 float64 //最大1.8e308
	fmt.Println(f1, f2)

	var ch byte
	ch = 'a'
	fmt.Println(ch)            //97
	fmt.Printf("c=%c\n", ch+1) //c=b
	ch1 := 97
	fmt.Printf("c=%c\n", ch1) //c=a

	//如果我们处理的字符是英文字符，使用byte，如果有中文，用rune
	var c1 rune = '慕'        //单引号字符，双引号是字符串
	fmt.Printf("c=%c\n", c1) //c=慕

	var name string = "i am tom"
	fmt.Printf("%s\n", name)
}
