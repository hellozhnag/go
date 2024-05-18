package main

import (
	"fmt"
	"time"
)

func main() {

	//条件判断和循环语句
	/*
		1、if条件判断
	*/
	a := 18
	country := "china"
	if a < 18 {
		fmt.Println("未成年~")
	} else if a == 18 && country == "china" {
		fmt.Println("刚满18岁~")
	} else {
	}

	/*
		2、for循环，go不提供while循环
	*/
	//这仨都不是必须的
	for i := 0; i < 10; i++ {
		fmt.Println(i) //打印0~9
	}
	//这种写法 无限循环
	//for {
	//	fmt.Println("hi~")
	//}

	var j int
	for j <= 3 {
		fmt.Println(j)
		j++
	}

	/*
		3、for循环打印99乘法表
	*/
	for a := 1; a <= 9; a++ {
		for b := 1; b <= a; b++ {
			fmt.Printf("%dx%d=%d ", a, b, a*b)
		}
		fmt.Println()
	}

	/*
		4、for range的循环用法
		主要是对字符串、数组、切片、map、channel
	*/
	name := "hi耶,go"
	//字符串的索引（key）、字符串对应的索引的字符值的拷贝（value）
	for key, value := range name {
		fmt.Println(key, value)
		fmt.Printf("%d,%c \r\n", key, value)
		//0 104
		//0,h
		//1 105
		//1,i
		//2 32822
		//2,耶
		//5 44
		//5,,
		//6 103
		//6,g
		//7 111
		//7,o
	}
	//其他写法
	for _, value := range name {
		fmt.Println(value)
	}
	// 不写key，返回索引
	for index := range name {
		fmt.Print(index)
		//打印对应index的字符
		fmt.Printf("%c ", name[index]) //只针对英文

		//0h 1i 2è 5, 6g 7o
	}
	fmt.Println()

	/*
		5、for循环的 break & continue
	*/
	test := 0
	for {
		time.Sleep(1 * time.Second)
		test++
		if test == 3 {
			continue
		}
		fmt.Print(test, " ") // 1 2 4 5
		if test == 5 {
			break
		}
	}

	/*
		6、goto语句
			很少用，可读性差，不建议用
			使用场景多用于错误处理，出错统一跳转到相应标签处，统一处理
	*/
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if j == 2 {
				//假设我们想在这结束两层循环，break只能结束一层
				goto over
			}
		}
	}
over:
	fmt.Println("over...")

	/*
		7、switch语句
			代码更可读、执行效率更高
	*/
	day := "星期三"
	switch day {
	case "星期一":
		fmt.Println("Monday...")
	case "星期二":
		fmt.Println("Tuesday...")
	case "星期三":
		fmt.Println("Wednesday...")
	case "星期四":
		fmt.Println("Thursday...")
	default:
		fmt.Println("other...")
	}

	score := 60
	switch {
	case score < 60:
		fmt.Println("不及格。。。")
	case score >= 60 && score < 80:
		fmt.Println("良好...")
	case score >= 80:
		fmt.Println("优秀")
	}

	switch score {
	case 60, 70, 80:
		fmt.Println("恰好是60 70 80.。。")
	default:
		fmt.Println("other...")
	}
}
