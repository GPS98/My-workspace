package main

import "fmt"

func main() {

	myword := "hello world"
	w1 := myword[0:5]
	w2 := myword[6:]
	fmt.Println(w1, "\n", w2)
}
