package main

import "fmt"

func rect(l int, b int) (area int, perimeter int) {
	area = l * b
	perimeter = 2 * (l + b)
	fmt.Println("area is ", area)
	fmt.Println("perimeter is ", perimeter)

	return
}

func main() {
	fmt.Println("enter length ")
	var l int
	fmt.Scanln(&l)
	fmt.Println("enter bredth ")
	var b int
	fmt.Scanln(&b)
	rect(l, b)

}
