package main

import "fmt"

// Funções com Retorno Nomeado
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main(){
	fmt.Println("Funções com Retorno Nomeado")

	// forma 1
	fmt.Println(calculosMatematicos(10, 20))

	// forma 2
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}