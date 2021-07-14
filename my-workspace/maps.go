package main

import "fmt"

func main() {

	company := make(map[string]int)
	//colour := make(map[string]int)
	company["ford"] = 12000
	company["tesla"] = 50000
	company["BMW"] = 42555
	company["Toyota"] = 20000
	company["Mercedes"] = 25000

	/*colour["red"] = 55
	colour["blue"] = 44
	colour["green"] = 33
	*/

	fmt.Println(company)

	company["Mazda"] = 15500
	fmt.Println("Updated:", company)

	delete(company, "Mercedes")
	fmt.Println("updated deleted:", company)

	for key, element := range company {
		fmt.Println(key, ">>", element)
	}

	/*carcompany := "tesla"
	cost := company[carcompany]

	fmt.Println("cost of ", carcompany, "is", cost)
	fmt.Println("cost of tesla", company["tesla"])*/

	value, ok := company["mercedes"]

	if ok == true {
		fmt.Println("cost of mercedes is", value)
		return
	}

	fmt.Println("notfound")
	fmt.Println(len(company))

}
