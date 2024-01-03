package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2
	return soma, subtracao, multiplicacao, divisao
}

func main (){
	soma := somar(10,20)
	fmt.Println(soma)

	// Sem retorno
	var f = func() {
		fmt.Println("Chamando a função f")
	}

	f()
	
	// Com retorno
	var txt = func(texto string) string {
		fmt.Println(texto)
		return texto
	}
	txt("Chamando a função txt com retorno")

	// Com vários (mais de 1) retornos
	soma, subtracao, multiplicacao, divisao := calculosMatematicos(10,20)
	fmt.Println(soma, " | ", subtracao, " | ", multiplicacao, " | ", divisao)

	// Se quiser ignorar um dos retornos
	_, _, resultMulti, _ := calculosMatematicos(10,20)
	fmt.Println(resultMulti)

}