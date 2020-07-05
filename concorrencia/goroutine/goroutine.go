package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d) \n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Gustavo", "Pq vc não fala comigo?", 3)
	// fale("Leticia", "Só posso falar depois de você!", 1)

	// go fale("Gustavo", "Ei...", 500)
	// go fale("Leticia", "Opa...", 500)
}
