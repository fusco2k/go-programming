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
		first:      "Miss",
		last:       "Fortune",
		iceFlavors: []string{"Strawberry", "Pineapple", "Nuts"},
	}

	y := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range y {
		fmt.Println(k)
		for k2, v2 := range v.iceFlavors {
			println(k2, v2)
		}
	}
}
