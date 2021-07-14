package main

import "fmt"

func main() {
	/*var exArray [5]string

	exArray[0] = "prem"
	exArray[1] = "prem"
	exArray[2] = "prem"
	exArray[3] = "prem"

	fmt.Println(exArray[0])
	fmt.Println(exArray[1])
	fmt.Println(exArray[2])
	fmt.Println(exArray[3])

	x := [6]int{0, 2, 4, 5, 2}

	fmt.Println(x)
	fmt.Println(len(x))

	z := [2]int{0: 25, 1: 65}
	fmt.Println(z)

	p := [5]int{1, 5, 7, 9, 4}
	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}

	for index, element := range p {
		fmt.Println(index, element)
	}*/

	/*d := [4]int{1, 2, 3, 4}
	g := d[1:3]
	fmt.Println(g)
	d[3] = 5
	fmt.Println(d)*/

	/*a := [...]int{2, 5, 4, 8, 9, 54, 7, 5, 9, 8, 4, 5, 4}
	for index, element := range a {
		fmt.Println("INdex:", index, "Element", element)
	}*/

	a := [5][3]int{
		{5, 6, 4},
		{7, 8, 9},
		{5, 9, 6},
		{2, 8, 6},
		{2, 2, 2}}

	var i, j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 3; j++ {
			fmt.Println("a[", i, "]", "b[&j]", a[i][j])
		}
	}
	fmt.Println(a)

}
