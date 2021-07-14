package main

import (
	"fmt"
	"time"
)

func test() {

	for i := 0; i < 10; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%d \n", i)
	}
}
func test1() {
	for b := 'a'; b < 'h'; b++ {
		time.Sleep(600 * time.Millisecond)
		fmt.Printf("%c \n", b)

	}

}

func main() {

	go test()
	go test1()
	time.Sleep(1 * time.Second)
	fmt.Println("exec Done")
}
