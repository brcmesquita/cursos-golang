package main

import "fmt"

func closure() func() {
	texto := "Dentro da função Closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	fmt.Println("Funções Closure")
	texto := "Dentro da função Main"
	fmt.Println(texto)
	
	funcaoNova := closure()
	funcaoNova()
}