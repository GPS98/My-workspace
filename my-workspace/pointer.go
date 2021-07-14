package main

import (
	"fmt"
)

func main() {
	var a int = 50
	var p *int

	p = &a

	fmt.Println("Value is ", p, &a, a)
}
