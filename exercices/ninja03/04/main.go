package main

import "fmt"

func main() {
	i := 1992
	for {
		if i <= 2019 {
			fmt.Println(i)
			i++
		} else {
			break
		}
	}
}
