package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int){
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 5, 5, 7 ,8 , 9, 133, 18899, 7833)
	fmt.Println(totalDaSoma)

	fmt.Println("------------")

	escrever("Olá mundo", 10, 20, 30, 40, 50, 55, 67)

}