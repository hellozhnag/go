package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	锁：解决对共享资源的竞争
*/

var total int32
var wg sync.WaitGroup
var lock sync.Mutex

// 锁能复制，但是复制之后就失去了锁的效果
//type Mutex struct {
//	state int32
//	sema  uint32
//}

func add() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(&total, 1)
		//lock.Lock()
		//total += 1
		//lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		//atomic为原子包
		atomic.AddInt32(&total, -1) //原子操作
		//lock.Lock()
		//total -= 1
		//lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
