package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(i ...int) int {
	var total int
	for _, v := range i {
		total += v
	}
	return total
}

func bar(i []int) int {
	var total int
	for _, v := range i {
		total += v
	}
	return total
}
