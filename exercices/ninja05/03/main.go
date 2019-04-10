package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	fordRanger := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	toyotaCorolla := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "silver",
		},
		luxury: false,
	}

	fmt.Println(fordRanger)
	fmt.Println(toyotaCorolla)
	fmt.Println(toyotaCorolla.color)
	fmt.Println(fordRanger.doors)
}
