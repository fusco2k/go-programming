package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`}}
	for k, v := range m {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println("\t", k2, v2)
		}
	}
	fmt.Println("//")
	m[`batman`] = []string{
		`Hero`, `Beaten up`, `Vodka`,
	}
	for k, v := range m {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println("\t", k2, v2)
		}
	}

	delete(m, `moneypenny_miss`)
	fmt.Println("//")
	for k, v := range m {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println("\t", k2, v2)
		}
	}
}
