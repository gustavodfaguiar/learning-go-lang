package main

import "fmt"

func imprimirDados(nome string, idade int) {
	fmt.Printf("%s tem %d anos. \n", nome, idade)
}

func soma(n, m int) int {
	return n + m
}

func main() {
	imprimirDados("Gustavo", 27)

	s := soma(3, 5)
	fmt.Println("A soma é", s)
}
