package main

import "fmt"

var x int
var y string
var z bool

var x2 int = 10
var y2 string = "Guguzinhu"
var z2 bool = false

type churrasco int

var costela churrasco
var picanha churrasco

func main() {
	value := 28
	name := "Gustavo Aguiar"
	dev := true

	fmt.Println("value: %v ", value)
	fmt.Println("name: %v ", name)
	fmt.Println("dev: %v ", dev)

	fmt.Println(x, y, z)

	s := fmt.Sprintln("%v %v %v", x2, y2, z2)

	fmt.Println(s)

	fmt.Println("valor: %v", x)
	fmt.Println("tipo: %T", x)

	x = 42

	fmt.Println(x)

	costela = 10

	fmt.Println(costela)

}
