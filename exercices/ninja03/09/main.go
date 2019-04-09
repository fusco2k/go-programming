package main

import "fmt"

func main() {
	favSport := "F1"
	switch favSport {
	case "football":
		fmt.Println("Chutar bola")
	case "F1":
		fmt.Println("This is not a Sport!")
	}
}
