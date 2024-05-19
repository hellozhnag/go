package main

import (
	"fmt"
	"time"
)

func asyncPrint() {
	fmt.Println("hi~")
}

// 主协程，主协程挂了，里面的子协程也要挂
func main() {

	/*
		协程基本使用
	*/
	//go语言没提供线程，只提供协程
	// 启动一个协程
	//go asyncPrint()
	//fmt.Println("hello world")
	//time.Sleep(1 * time.Second)

	// 输出：
	//hello world
	//hi~

	/*
		go 匿名函数
	*/
	//go func() {
	//	for {
	//		time.Sleep(1 * time.Second)
	//		fmt.Println("hi~")
	//	}
	//}()
	//time.Sleep(10 * time.Second)
	// 过一秒打印一次hi~，直到10s后结束

	/*
		启动多个上百个协程
	*/
	//1、闭包，匿名函数访问外部函数的局部变量
	//2、for循环：每次for循环，i变量会被重用
	// goroutine每次运行时，里面得i其实指向的是外部i，并非复制
	//所以这里的输出无法控制
	//for i := 0; i < 1000; i++ {
	//	go func() {
	//		for {
	//			fmt.Println(i)
	//		}
	//	}()
	//}

	// 所以我们要修改成这样：
	// 这个就运行正常了
	for i := 0; i < 100; i++ {
		tmp := i
		go func() {
			fmt.Println(tmp) // goroutine运行顺序不确定
		}()
	}
	// 或者这样(函数值传递）
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	//这个地方到底要sleep多久？
	time.Sleep(1010 * time.Second)
	// 是否可不写sleep，使的所有goroutine都结束后，主协程也自动结束-waitgroup

}
