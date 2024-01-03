package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	fmt.Println("---------------------")
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	fmt.Println("---------------------")
	variavel2++
	fmt.Println(variavel1, variavel2)

	fmt.Println("---------------------")
	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // Precisa usar o & senão não não compila

	fmt.Println(variavel3, ponteiro) // 100 0xc000014110
	fmt.Println(variavel3, *ponteiro) // 100 100

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // 150 150

}