package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO

	// CONCORRÊNCIA é quando duas coisas rodam uma após a outra, dando a impressão
	// de que estão rodando juntas. Roda um pouco de uma, depois um pouco de outra
	// até rodar todos os processos.

	// PARALELISMO é quando duas coisas rodam juntas, ao mesmo tempo, uma em cada
	// núcleo do processador. Uma não depende da outra pra rodar.

	go escrever("Olá mundo!") // goroutine
	escrever("Programando em Go!")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
