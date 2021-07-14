package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch today.Day() {
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
