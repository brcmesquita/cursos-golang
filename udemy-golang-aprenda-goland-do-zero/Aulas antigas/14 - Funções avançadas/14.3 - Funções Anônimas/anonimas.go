package main

import "fmt"

func main() {

	// Função anônima no Go
	func() {
		fmt.Println("Olá mundo!")
	}() // para chamar uma função anônima, usamos parênteses

	// Funções anônimas com parâmtros
	func(texto string) {
		fmt.Println(texto)
	}("Olá mundo com parâmetros")

	// Funções anônimas com parâmetros e retorno
	retorno := func(texto string) string {
		return fmt.Sprintf("Seja bem-vindo %s", texto)
	}("Bruno Raphael")
	fmt.Println(retorno)
}