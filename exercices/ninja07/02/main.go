package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	(*p).first = "James"
}
func main() {
	p1 := person{
		first: "Bruce",
		last:  "Wayne",
		age:   26,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
