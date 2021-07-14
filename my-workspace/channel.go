package main

import (
	"fmt"
	"time"
)

/*
func add(a int, b int, plus chan int) {

	total := 0
	total = a + b
	plus <- total
}

func sub(a int, b int, minus chan int) {

	total := 0
	total = a - b
	minus <- total
}

func main() {
	a := 54
	b := 35
	additionch := make(chan int)
	substractionch := make(chan int)
	go add(a, b, additionch)
	go sub(a, b, substractionch)

	additions, substractions := <-additionch, <-substractionch

	fmt.Println("total is", additions+substractions)

}

*/

func name(personname chan string) {
	fmt.Println("Enter your First Name")
	var fname string
	fmt.Scanln(&fname)
	fmt.Println("Enter your Last Name")
	var lname string
	fmt.Scanln(&lname)
	fullname := fname + lname
	personname <- fullname

}

func age(personage chan int) {
	fmt.Println("enter your age")
	var prage int
	fmt.Scanln(&prage)
	totalAge := prage
	personage <- totalAge
}

func main() {
	namedetailsch := make(chan string)
	agedetailsch := make(chan int)

	go name(namedetailsch)
	time.Sleep(5 * time.Second)
	go age(agedetailsch)

	ndetails, adetails := <-namedetailsch, <-agedetailsch

	fmt.Println("total name and age ", ndetails, adetails)
}
