package main

import "fmt"

type person struct{
	name string
	age int
}

func (p *person) speak(){
	fmt.Println("Hello, I'm", p.name)
}

type human interface {
	speak()
}

func saySomething (h human){
	h.speak()
}

func main() {
	p1 := person{
		name: "James",
		age:36,
	}

	saySomething(&p1)
	//saySomething(p1)
}
