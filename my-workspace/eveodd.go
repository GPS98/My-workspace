package main

import (
	"fmt"
)

func main() {
	/*x := 100
	if x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Canada")
	}*/
	//########################
	k := 1

	for k = 0; k < 10; k++ {

		if k%2 == 0 {
			fmt.Println("even", k)
		} else {
			fmt.Println("odd", k)
		}
	}
//############################################
}

