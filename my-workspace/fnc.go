package main

import "fmt"

func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println("total is", total)

}

func main() {
	add(30, 50)
}
