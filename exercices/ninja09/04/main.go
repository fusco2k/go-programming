package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var ws sync.WaitGroup
	var mx sync.Mutex

	ws.Add(2)
	fmt.Println("Counter at start:", counter)
	go func() {
		for i := 0; i < 100; i++ {
			mx.Lock()
			z := counter
			z++
			counter = z
			fmt.Println("Counter at first routine:", counter)
			mx.Unlock()
		}
		ws.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			mx.Lock()
			z := counter
			z++
			counter = z
			fmt.Println("Counter at second routine:", counter)
			mx.Unlock()
		}
		ws.Done()
	}()
	ws.Wait()
	fmt.Println("Counter at end:", counter)
}
