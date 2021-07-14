package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
}
type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func measure(g geometry) {
	fmt.Println(g.area())

}
func main() {
	c := circle{radius: 6}
	r := rectangle{width: 4, height: 7}

	measure(c)
	measure(r)

}
