package main

import (
	"fmt"
	"time"
)

func numbers(num chan int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets(alp chan string) {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	numbo := make(chan int)
	alpho := make(chan string)

	go numbers(numbo)
	go alphabets(alpho)
	time.Sleep(10 * time.Second)
	numberout, alphaout := numbo, alpho
	fmt.Println("main terminated", numberout, alphaout)
}
