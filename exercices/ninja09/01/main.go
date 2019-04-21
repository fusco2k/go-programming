package main

import (
	"fmt"
	"sync"
)

func main() {

	var ws sync.WaitGroup
	ws.Add(2)

	go func() {
		fmt.Println("first routine")
		ws.Done()
	}()

	go func() {
		fmt.Println("second routine")
		ws.Done()
	}()
	fmt.Println("Waiting to exit")
	ws.Wait()
}
