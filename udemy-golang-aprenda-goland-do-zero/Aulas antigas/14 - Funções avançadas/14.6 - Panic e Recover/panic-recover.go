package main

import "fmt"

// vamos criar um recover para manter o programa funcionando
func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar a execução!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!") // Sem um recover, o programa morre
}

func main() {
	fmt.Println("Panic e Recover")
	fmt.Println("---------------")
	
	fmt.Println(alunoEstaAprovado(6,7))
	fmt.Println("Pós execução")
}