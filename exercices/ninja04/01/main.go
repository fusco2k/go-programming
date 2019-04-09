package main

import "fmt"

func main() {
	x := [5]int{23, 24, 25, 26, 27}

	for i, v := range x {
		fmt.Printf("%v\t%v\n", i, v)
	}
	fmt.Printf("%T\n", x)
}
