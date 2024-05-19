package main

import (
	"fmt"
	"time"
)

// goroutine间通信采用channel

func main() {
	/*
		go理念：
			不要通过共享内存来通信
			而要通过通信来共享内存

		像java提供了消息队列机制-生产者消费者之间的关系
		channel再加上语法糖让使用channel更加简单
	*/
	var msg chan string        //不初始化的话为nil
	msg = make(chan string, 1) //如果初始化值的0的话，会死锁，永远阻塞
	msg <- "bobby"             //表示放值到channel中
	data := <-msg              //从channel中取值
	fmt.Println(data)          //bobby

	// 消费无缓冲的数据,
	var msg2 chan string
	msg2 = make(chan string, 0) //无缓冲的channel
	// 无缓冲channel要启动一个goroutine去消费他
	go func(msg chan string) { // go有一种happen-before的机制
		data := <-msg
		fmt.Println(data)
	}(msg2)
	msg2 <- "bobby" // 放值到channel中

	//出现deadlock
	//	- waitgroup少了done调用
	//	- 无缓冲的channel

	/*
		有缓冲channel、无缓冲channel的使用场景
			- 无缓冲channel 适用于通知，A一旦发出消息后B立马执行
			- 有缓冲channel 适用于消费者生产者之间的通信
		go中channel应用场景
			- 消息传递  消息过滤
			- 信号广播
			- 事件订阅和广播
			- 任务分发
			- 结果汇总
			- 并发控制
			- 同步异步
			- ...
	*/

	/*
		forrange 对channel遍历
			- 让消费者源源不断取数据
	*/
	var msg3 chan string
	msg3 = make(chan string, 1)
	go func(msg chan string) {
		for data := range msg3 {
			fmt.Println(data) //把值都打印出来,会一直等待值，并不会打印下面all done
		}
		// 如何让通知for停止，close（msg）
		fmt.Println("all done 。。。")
		//打印：
		//bobby1
		//bobby2
		//bobby3
		//all done 。。。
	}(msg3)
	msg3 <- "bobby1"
	msg3 <- "bobby2"
	msg3 <- "bobby3"
	close(msg3)
	// 已经关闭的channel不能再放值了，不然报错panic: send on closed channel
	// 但是可以继续取值
	//msg3 <- "bobby4"

	/*
		单向channel
			- channel默认双向：收发
			- 但是通常channel进行参数传递，不希望对方写
	*/
	//var ch1 chan int       //双向channel
	//var ch2 chan<- float64 //单向channel，只能写入float64
	//var ch3 <-chan int     //单向channel，只能读取

	c := make(chan int, 3)
	var send chan<- int = c // send-only
	var read <-chan int = c // read-only
	send <- 1
	<-read
	time.Sleep(time.Second * 1)
	//<-send //报错
	//d1 := (chan int)(send)//不能将单向channel转换成双向channel
	go func(out chan<- int) {
		for i := 0; i < 10; i++ {
			out <- i * i
		}
	}(c) //会自动转换
	go func(in <-chan int) {
		for data := range in {
			fmt.Printf("%d ", data)
			//打印：
			//0 1 4 9 16 25 36 49 64 81
		}
		fmt.Println()
	}(c)

	/*
		通过channel实现交叉打印
			- 两个goroutine，一个打印数字，一个打印字母，交叉打印
				- 先打印数字
	*/
	time.Sleep(time.Second * 1)
	var number, letter = make(chan bool), make(chan bool)
	go func() {
		i := 1
		for {
			<-number //等待通知
			fmt.Printf("%d", i)
			i += 2
			letter <- true
		}
	}()

	go func() {
		i := 1
		str := "ABC"
		for {
			<-letter
			fmt.Print(str)
			i += 2
			number <- true
		}
	}()
	number <- true //先打印数字
	time.Sleep(time.Second * 2)

}
