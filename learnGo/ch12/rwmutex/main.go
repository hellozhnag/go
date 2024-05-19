package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁
// 锁让并行代码串行化
// 即使设计锁，也尽量保证并行

// 有两组协程，一组负责读，另一组负责读

func main() {
	//var num int
	var rwlock sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(6)

	//写的goroutine
	go func() {
		time.Sleep(1 * time.Second)
		defer wg.Done()
		// 加写锁，防止写锁、读锁获取
		rwlock.Lock()
		defer rwlock.Unlock()
		fmt.Println("get write lock")
		time.Sleep(5 * time.Second)
	}()

	//多个goroutine
	for i := 0; i < 5; i++ {
		//读的goroutine
		go func() {
			defer wg.Done()
			for {
				//加读锁，不会阻止读锁
				rwlock.RLock()
				time.Sleep(500 * time.Millisecond)
				fmt.Println("get read lock")
				rwlock.RUnlock()
			}
		}()
	}
	wg.Wait()

}
