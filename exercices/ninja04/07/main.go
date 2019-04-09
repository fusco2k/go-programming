package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xy := [][]string{x, y}
	for i, c := range xy {
		fmt.Println(i, c)
		for i2, c2 := range c {
			fmt.Printf("0%v\t%v\n", i2, c2)
		}
	}
}
