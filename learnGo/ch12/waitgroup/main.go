package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	//监控多少个goroutine结束
	//wg.Add(100) // 或者写在for中
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			////Add方法 要和 Done方法成对出现,Done表示goroutine结束
			//wg.Done() //我们一般放在第一行 并用defer
		}(i)
	}
	wg.Wait()
}
