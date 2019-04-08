package main

import "fmt"

var x = 2

func main() {
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#U\n", x)
	y := x << 1
	fmt.Printf("%d\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%#U\n", y)
}
