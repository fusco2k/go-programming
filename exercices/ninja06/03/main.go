package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hey there!")
}

func foo() {
	fmt.Println("Foo here!")
}
