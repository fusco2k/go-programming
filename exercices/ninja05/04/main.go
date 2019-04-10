package main

import "fmt"

func main() {
	filipe := struct {
		name   string
		age    int
		height int
	}{
		name:   "Filipe Fusco",
		age:    26,
		height: 167,
	}

	fmt.Println(filipe)
}
