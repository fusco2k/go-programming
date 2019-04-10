package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (person) speak(p person) {
	fmt.Println("My name is", p.first, p.last)
	fmt.Println("My age is", p.age)
}

func main() {
	p1 := person{
		first: "Bruce",
		last:  "Wayne",
		age:   36,
	}

	p1.speak(p1)
}
