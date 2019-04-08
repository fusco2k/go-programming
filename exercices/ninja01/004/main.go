package main

import "fmt"

type cube int
var x cube

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)
}