package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	var ws sync.WaitGroup
	ws.Add(2)
	fmt.Println("Counter at start:", counter)
	go func() {
		for i := 0; i < 100; i++ {
			z := counter
			runtime.Gosched()
			z++
			counter = z
			fmt.Println("Counter at first routine:", counter)
		}
		ws.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			z := counter
			runtime.Gosched()
			z++
			counter = z
			fmt.Println("Counter at second routine:", counter)
		}
		ws.Done()
	}()
	ws.Wait()
	fmt.Println("Counter at end:", counter)
}
