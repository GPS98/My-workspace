package main

import (
	"fmt"
	"learnpackage/simpleinterest"
)

func main() {
	a, b := 50, 10
	total := simpleinterest.Add(a, b)
	fmt.Println(" total is ", total)
}

/*{
	fmt.Println("Simple interest calculation")
	p := 5000.0
	r := 10.0
	t := 1.0
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}*/
