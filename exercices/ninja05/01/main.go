package main

import "fmt"

type person struct {
	first      string
	last       string
	iceFlavors []string
}

func main() {
	p1 := person{
		first:      "Bruce",
		last:       "Wayne",
		iceFlavors: []string{"Vanilla", "Coffee", "Chocolate"},
	}
	p2 := person{
		first: "Miss",
		last:  "Fortune",
		s:     []string{"Strawberry", "Pineapple", "Nuts"},
	}

	y := []person{p1, p2}

	for _, v := range y {
		fmt.Println(v.first, v.last)
		for k2, v2 := range v.iceFlavors {
			println(k2, v2)
		}
	}
}
