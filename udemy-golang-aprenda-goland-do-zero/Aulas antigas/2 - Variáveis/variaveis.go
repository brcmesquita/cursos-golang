package main

import "fmt"

func main() {
	// Declarando variáveis

	// Modo 1
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	// Modo 2 - Inferência de Tipo
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	// Declarando várias variáveis ao mesmo tempo
	// Modo 1
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)
	fmt.Println(variavel3)
	fmt.Println(variavel4)

	// Modo 2
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5)
	fmt.Println(variavel6)

	// Declarando Constantes
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// Trocando valores entre variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}