package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// Sequência de Fibonacci - 1 1 2 3 5 8 13
	posicao := uint(133)
	fmt.Println(fibonacci(posicao))
}