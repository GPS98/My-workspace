package main

import "fmt"

type car struct {
	name    string
	engine  string
	mailage int
	colour  string
}

func main() {
	var car1 car
	var car2 car

	car1.name = "elantra"
	car1.engine = "v4"
	car1.mailage = 12548
	car1.colour = "Black"

	car2.name = "mustang"
	car2.engine = "v8"
	car2.mailage = 15487
	car2.colour = "Red"

	fmt.Println(&car1)
	fmt.Println(&car2)

	/*fmt.Println("first car details are ", car1.name, car1.engine, car1.mailage, car1.colour)
	fmt.Println("second car details are ", car2.name, car2.engine, car2.mailage, car2.colour)
	*/
}
