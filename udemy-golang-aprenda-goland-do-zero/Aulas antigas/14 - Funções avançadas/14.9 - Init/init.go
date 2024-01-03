package main

import "fmt"

var n int

// A função INIT é sempre executada primeiro, mesmo estando no final do arquivo
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")

	fmt.Println(n)
}