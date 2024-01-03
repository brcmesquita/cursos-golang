package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa // não precisa declarar um tipo
	curso string
	faculdade string
}

func main() {
fmt.Println("Herança")

p1 := pessoa{"Bruno", "Raphael", 37, 177}
fmt.Println(p1)

e1 := estudante{p1, "Sistemas de Informação", "Estácio de Sá"}
fmt.Println(e1)
fmt.Println(e1.nome, e1.sobrenome)
}