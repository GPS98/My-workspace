package main

import "fmt"

type sum struct {
	a, b int
}

func (s sum) addition() int {
	return s.a + s.b
}

func main() {
	s := sum{a: 54, b: 87}
	fmt.Println(s.addition())
}

/*type Circle struct {
	x, y, radius float64
}

func (c Circle) calc() float64 {

	return c.x * c.y * c.radius

}

func main() {
	c := Circle{x: 10, y: 10, radius: 12}
	fmt.Println(c.calc())
}
*/
