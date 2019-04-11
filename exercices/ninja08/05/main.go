package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []Person

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := Person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := Person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []Person{u1, u2, u3}

	fmt.Println(users)

	for _, v := range users {
		sort.Strings(v.Sayings)
	}
	fmt.Println("By Name: ")
	sort.Sort(ByLast(users))

	for _, v := range users {
		fmt.Println(v.First, v.Last, "at age", v.Age)
		fmt.Println("\t Quotes: ")
		for k, v2 := range v.Sayings {
			fmt.Println("\t", k, v2)
		}
	}

	fmt.Println("By Age: ")
	sort.Sort(ByAge(users))

	for _, v := range users {
		fmt.Println(v.First, v.Last, "at age", v.Age)
		fmt.Println("\t Quotes: ")
		for k, v2 := range v.Sayings {
			fmt.Println("\t", k, v2)
		}
	}
}
