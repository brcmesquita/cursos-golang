package main

import "fmt"

func main(){
	fmt.Println("Estruturas de controle")

	numero := 10

	// Parênteses não são obrigatórios
	if numero > 15 {
		fmt.Println("Maior que 15")
	}

	// Não funciona sem as chaves
	// if numero > 15
	//   fmt.Println("Maior que 15")

	// Não funciona sem as chaves na mesma linha
	// if numero > 15 fmt.Println("Maior que 15")

	// Mas funciona com as chaves na mesma linha
	if numero > 15 { fmt.Println("Maior que 15") }

	// Exemplo com If e Else
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// Também podemos fazer desta maneira
	// atribuindo uma nova variável e atribuindo um valor
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else {
		fmt.Println("Número é menor que zero")
	}
	// a variável não existe fora do escopo
	// a linha abaixo não funcionará
	//fmt.Println(outroNumero)

	// Também podemos fazer desta outra maneira
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que - 10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}