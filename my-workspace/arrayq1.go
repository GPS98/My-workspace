package main

import "fmt"

func main() {

	/*a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(a)*/

	names := [5]string{"sagar", "prem", "ford", "bmw"}
	fmt.Println(names)
	fmt.Println(len(names), cap(names))

}
