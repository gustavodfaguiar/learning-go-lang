package main

import "fmt"

func main() {
	boolean := true

	if boolean {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	number := 2

	switch number {
	case 2:
		fmt.Println("Number 2")
	default:
		fmt.Println("Default")
	}

}
