package main

import "fmt"

type house struct {
	noRoom string
	price  string
	city   string
}

func main() {

	var details house
	var details1 house

	details.noRoom = "available"
	details.price = "1000"
	details.city = "avenel"

	details1.noRoom = "not available"
	details1.price = "100000"
	details1.city = "cape"

	fmt.Println("room one details:", details)
	fmt.Println("room two details", details1)

}
