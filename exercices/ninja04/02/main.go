package main

import "fmt"

func main() {
	x := []int{22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	for i, c := range x {
		fmt.Printf("%v\t%v\n", i, c)
	}
	fmt.Printf("\n%T\n", x)
}
