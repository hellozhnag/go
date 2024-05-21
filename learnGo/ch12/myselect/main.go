package main

import (
	"fmt"
	"time"
)

//var done = make(chan struct{}) //空结构体0空间，省内存
// channel是协程安全的

func g1(ch chan struct{}) {
	time.Sleep(time.Second)
	ch <- struct{}{}
}
func g2(ch chan struct{}) {
	time.Sleep(time.Second * 2)
	ch <- struct{}{}
}
func main() {
	// select 写法上类似于switch case语句
	// 但是功能上和linux提供的io的select、poll、epoll相似
	// select主要作用与多个channel
	//		- 有两个goroutine在执行，当一个结束后，主goroutine立马哪个goroutine结束了

	g1Ch, g2ch := make(chan struct{}, 1), make(chan struct{}, 1)
	go g1(g1Ch)
	go g2(g2ch)
	//某个分支就绪了就执行该分支
	//如果俩都刚好就绪了，随机执行;防止饥饿，如果在前边的先执行，那么若每次g1和g2都同时就绪，那么g2永远不会执行

	//防止因g1或者g2阻塞而导致整个函数阻塞，我们设置一个超时定时器，1s
	tc := time.NewTicker(time.Second)
	select {
	case <-g1Ch:
		fmt.Println("g1 done ...")
	case <-g2ch:
		fmt.Println("g2 done ...")
	case <-tc.C:
		fmt.Println("default ....")
		return
	}

	time.Sleep(time.Second * 3)

}
