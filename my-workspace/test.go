package main

import "fmt"

func main() {

	fmt.Println("enter number")
	var n int
	fmt.Scanln(&n)

	switch n {
	case 5, 10:
		fmt.Println("good day")
		fallthrough
	case 12, 15:
		fmt.Println("wonderfulday")
		fallthrough
	case 22, 25:
		fmt.Println("second day")
		fallthrough
	default:
		fmt.Println("nothing")
	}

}
