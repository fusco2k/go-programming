package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for z := 0; z < 3; z++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
