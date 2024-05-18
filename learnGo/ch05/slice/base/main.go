package main

import (
	"fmt"
	"strconv"
)

func printSlice(data []string) {
	data[0] = "java"
	for i := 0; i < 10; i++ {
		data = append(data, strconv.Itoa(i))
	}
	fmt.Println(data)
}

func main() {
	//slice 底层原理

	//对于前边内容的一个补充
	fmt.Printf("%T, %T\n", [...]string{}, []string{}) //[0]string, []string

	/*
		go的slice在函数参数传递的时候是值传递，
		但是效果又呈现出了引用的效果（不完全是）
	*/
	//现象引入：
	course := []string{"go", "grpc", "gin"}
	printSlice(course)  ////[java grpc gin 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(course) //[java grpc gin]

	/*
		type slice mystruct {
			array unsafe.Pointer //用来存储实际数据的数组指针，指向一块连续的内存
			len   int            //切片中元素数量
			cap   int            //array数组的长度
		}
	*/

	// 原理：

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := data[3:6]
	s2 := data[2:7]
	s2[1] = 22
	fmt.Println(s1)               //[22 5 6]
	fmt.Println(s2)               //[3 22 5 6 7]
	fmt.Println(len(s1), cap(s1)) //3 7
	fmt.Println(len(s2), cap(s2)) //5,8

	s2 = append(s2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1, 1)
	s2[1] = 33
	fmt.Println(s1)               //[22 5 6]，扩容后就不影响了
	fmt.Println(s2)               //[3 33 5 6 7 1 2 3 4 5 6 7 8 9 1 1 1]
	fmt.Println(len(s1), cap(s1)) //3 7
	fmt.Println(len(s2), cap(s2)) //17 18

	var data1 []int
	for i := 0; i < 2000; i++ {
		data1 = append(data1, i)
		fmt.Printf("len: %d, cap: %d\n", len(data1), cap(data1))
	}
}
