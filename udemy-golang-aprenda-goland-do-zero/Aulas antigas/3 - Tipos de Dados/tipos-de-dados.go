package main

import (
	"errors"
	"fmt"
)

func main(){
	// Tipos Inteiros -> int8, int16, int32, int64
	var numero int8 = 100
	fmt.Println(numero)

	// Tipos Inteiros Unsigned -> uint8, uint16, uint32, uint64
	var numero2 uint8 = 100
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// alias
	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Tipos Reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 678901234567890.12
	fmt.Println(numeroReal2)

	// Float com Inferência
	numeroReal3 := 123456789.0
	fmt.Println(numeroReal3)

	// FIM NÚMEROS

	// Tipos String

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	// Não existe CHAR no Go

	// FIM STRINGS

	// Falando sobre valor 0 -> Sempre que declara uma variável e não atribui valor

	var texto0 string
	fmt.Println(texto0) // é uma string vazia, pois não foi atribuído nenhum valor

	var numero0 int16
	fmt.Println(numero0) // vai imprimir 0

	var numeroReal0 float32
	fmt.Println(numeroReal0) // vai imprimir 0

	// TIPO BOOLEANO
	var booleano1 bool = true // ou false
	fmt.Println(booleano1)
	// false é igual a 0 (zero)

	// TIPO ERRO
	var erro error
	fmt.Println(erro) // <nil>

	// PARA CRIAR UM TIPO DE ERRO
	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}