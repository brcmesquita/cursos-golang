package main

import "fmt"

// o DEFER faz com que algo (que ele aponta) seja executado por último no escopo
func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado...")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcao1() // defer = adiar
	funcao2()
	fmt.Println(alunoEstaAprovado(7, 8))
}