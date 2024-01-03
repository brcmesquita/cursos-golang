package main

import "fmt"

// No Switch, o Go não tem o parâmetro "break", ou seja,
// assim que ele encontrar uma condição verdadeira,
// ele executa a condição e depois sai.

// forma 1
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

// forma 2
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta"
	case numero == 6:
		return "Sexta"
	case numero == 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

// forma 3
func diaDaSemana3(numero int) string {
 	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

// forma 4 - Fallthrough
func diaDaSemana4(numero int) string {
	var diaDaSemana string

 switch {
 case numero == 1:
	 diaDaSemana = "Domingo"
	 fallthrough // se o resultado for o de cima, ele pula para o próximo ou algo assim
 case numero == 2:
	 diaDaSemana = "Segunda"
 case numero == 3:
	 diaDaSemana = "Terça"
 case numero == 4:
	 diaDaSemana = "Quarta"
 case numero == 5:
	 diaDaSemana = "Quinta"
 case numero == 6:
	 diaDaSemana = "Sexta"
 case numero == 7:
	 diaDaSemana = "Sábado"
 default:
	 diaDaSemana = "Número inválido"
 }
 return diaDaSemana
}

func main(){
	fmt.Println("Switch")

	dia := diaDaSemana(10)
	fmt.Println(dia)

	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)

	dia3 := diaDaSemana3(4)
	fmt.Println(dia3)
	
	dia4 := diaDaSemana4(1)
	fmt.Println(dia4)
}