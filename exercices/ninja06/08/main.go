package main

import "fmt"

func main() {
	x := foo()
	x()
}

func foo() func() {
	return func() {
		fmt.Println("Hello")
	}
}
