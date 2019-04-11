package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var ws sync.WaitGroup

	ws.Add(2)
	fmt.Println("Counter at start:", counter)
	go func() {
		for i := 0; i < 100; i++ {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter at first routine:", atomic.LoadInt64(&counter))
		}
		ws.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter at second routine:", atomic.LoadInt64(&counter))
		}
		ws.Done()
	}()
	ws.Wait()
	fmt.Println("Counter at end:", counter)
}
