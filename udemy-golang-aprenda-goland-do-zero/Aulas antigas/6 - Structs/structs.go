package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main()  {
	fmt.Println("Arquivos structs")

	// modo 1
	var u usuario
	u.nome = "Bruno"
	u.idade = 37
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua ABC", 123}

	// modo 2
	usuario2 := usuario{"Stéphanie", 36, enderecoExemplo}
	fmt.Println(usuario2)

	// modo 3 (quando faltam dados)
	usuario3 := usuario{nome: "Bruno Raphael"}
	fmt.Println(usuario3)
}