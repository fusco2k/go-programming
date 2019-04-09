package main

import "fmt"

func main() {
	x := 41
	if x == 42 {
		fmt.Println("Meaning of life")
	} else if x != 42 {
		fmt.Println("I'm lost")
	}
}
