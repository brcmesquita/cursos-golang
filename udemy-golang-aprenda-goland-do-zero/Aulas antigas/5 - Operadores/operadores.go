package main

import "fmt"

func main(){
	// ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 10 * 5
	divisao := 10 / 4
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	// Não é possível fazer operações com tipos de dados diferentes
	var numero1 int16 = 10
	var numero2 int32 = 25
	//soma2 := numero1 + numero2 // vai dar erro

	// primeira solução seria trocar os tipos
	// a segunda é fazer um parse
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)

	// por segurança, fazemos um parse no tipo menor para o tipo maior
	soma3 := int32(numero1) + numero2
	fmt.Println(soma3)

	// FIM DOS ARITMETICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String 1"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)
	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 < 2) // 1 menor que 2
	fmt.Println(1 > 2) // 1 maior que 2
	fmt.Println(1 <= 2) // 1 menor ou igual a 2
	fmt.Println(1 >= 2) // 1 maior ou igual a 2
	fmt.Println(1 == 2) // 1 igual a 2
	fmt.Println(1 != 2) // 1 diferente de 2
	// FIM DOS RELACIONAIS

// OPERADORES LÓGICOS
verdadeiro, falso := true, false
fmt.Println(verdadeiro && falso) // E
fmt.Println(verdadeiro || falso) // OU
fmt.Println(!verdadeiro) // Negação de true, será igual a false
fmt.Println(!falso) // Negação de false, será igual a true
// FIM DOS OPERADORES LÓGICOS

// OPERADORES UNÁRIOS
numero := 10
numero++ // aumentar o valor em 1
numero-- // decrementa o número em 1
numero += 15 // igual a numero = numero + 15
numero -= 15 // igual a numero = numero - 15
numero *= 15 // igual a numero = numero * 15
numero /= 15 // igual a numero = numero / 15
numero %= 15 // igual a numero = numero % 15

fmt.Println(numero)
// FIM DOS OPERADORES UNÁRIOS

// OPERADORES TERNÁRIOS
// NÃO EXISTEM OPERADORES TERNÁRIOS NO GO
// texto := numero > 5 ? "Número é maior do que 5" : "Número é menor do que 5"

// COMO QUE RESOLVE? USANDO IF NORMAL
var texto string
if numero > 5 {
	texto = "Maior que 5"
} else {
	texto = "Menor que 5"
}
fmt.Println(texto)

}