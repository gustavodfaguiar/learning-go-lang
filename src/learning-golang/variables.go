package main

import "fmt"

var name_test = "Test"

func main() {

	print(name_test)

	name := "Gustavo"
	fmt.Print(name)

	const name_two = "Gustavo"
	fmt.Print(name_two)

	number1, number2 := 10, 2

	mult := number1 * number2
	fmt.Print(mult)

}
