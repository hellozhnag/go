package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{}) //空结构体0空间，省内存
// channel是协程安全的

func g1() {
	time.Sleep(time.Second)
	fmt.Println("g1")
	done <- struct{}{}
}
func g2() {
	time.Sleep(time.Second * 2)
	fmt.Println("g2")
	done <- struct{}{}
}
func main() {
	// select 写法上类似于switch case语句
	// 但是功能上和linux提供的io的select、poll、epoll相似
	// select主要作用与多个channel
	//		- 有两个goroutine在执行，当一个结束后，主goroutine立马哪个goroutine结束了

	go g1()
	go g2()

	<-done
}
