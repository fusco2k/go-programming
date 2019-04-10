package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(ii...))
	fmt.Println(odd(sum, ii...))
}

func sum(i ...int) int {
	var sum int
	for _, v := range i {
		sum += v
	}
	return sum
}

func odd(f func(i ...int) int, ii ...int) int {
	var yy []int
	for _, v := range ii {
		if v%2 == 1 {
			yy = append(yy, v)
		}
	}
	return f(yy...)
}
