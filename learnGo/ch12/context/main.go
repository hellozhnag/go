package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cpuInfo(stop chan struct{}) {
	defer wg.Done()
	for true {
		select {
		//实现父goroutine向子goroutine传递信息
		case <-stop:
			fmt.Println("退出CPU 监控程序")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("CPU信息~")
		}
	}
}

// context是一个接口，可以传任意参数
func cpuInfo1(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出CPU监控~")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("CPU信息~")
		}
	}
}
func cpuInfo2(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出CPU监控~")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("CPU信息~")
		}
	}
}
func cpuInfo3(ctx context.Context) {
	defer wg.Done()
	fmt.Printf("nihao:%s\n", ctx.Value("nihao"))
	//我们可以记录日志，浙醋请求是来自哪个key
}

func main() {
	//有一个goroutine监控CPU信息,每两秒打印下消息
	wg.Add(1)
	//stop := make(chan struct{})
	//go cpuInfo(stop)

	//6秒后告诉子goroutine停止
	//time.Sleep(time.Second * 6)
	//stop <- struct{}{}
	//wg.Wait()

	/*
		使用context更加优雅~
			- context包提供了3种函数
				- WithCancel---主动取消
				- WithTimeout---主动超时
				- WithValue
	*/
	//WithCancel生成一个子context
	// 如果goroutine，函数希望被控制（超时、传值。。。）
	// 但是不希望影响原来的接口信息，那么函数的第一个参数为ctx
	ctx, cancel := context.WithCancel(context.Background())
	// 如果，传递的是子ctx，也行~
	ctx2, _ := context.WithCancel(ctx)
	go cpuInfo1(ctx2)
	//go cpuInfo1(ctx)
	time.Sleep(time.Second * 6)
	cancel()

	wg.Add(1)
	//withtimeout---内部是withdeadline
	//设置6s超时时间
	ctx3, _ := context.WithTimeout(context.Background(), time.Second*6)
	go cpuInfo2(ctx3)

	// withvalue
	wg.Add(1)
	ctx4 := context.WithValue(ctx3, "nihao", "hello~")
	go cpuInfo3(ctx4)
	wg.Wait()
	fmt.Println("监控完成~")
}
