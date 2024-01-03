package main

import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Println("Arrays e Slices") // Arrays e Slices

	// Criando um Array em Go
	// Os arrays são engessados, eles possuem tamanhos estáticos, que precisam ser declarados
	// Cláusula Var + Nome do Array + Tamanho do Array + Tipo de dados que o Array irá conter
	var array1 [5]int
	fmt.Println(array1) // [0 0 0 0 0]

	// Como adicionar elementos ao Array
	array1[0] = 1
	fmt.Println(array1) // [1 0 0 0 0]

	// modo 2
	array2 := [5]string{"Bruno", "Raphael", "Cabral", "de", "Mesquita"}
	fmt.Println(array2) // [Bruno Raphael Cabral de Mesquita]

	// modo 3 - não é muito utilizado
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// FIM ARRAY

	// SLICE
	// O slice não é engessado como o Array, pois pode ter tamanho dinâmico, não sendo necessário declarar seu tamanho
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice) // [10 11 12 13 14 15 16 17]

	// reflect.TypeOf() - retorna o tipo da variável
	fmt.Println(reflect.TypeOf(slice)) // []int
	fmt.Println(reflect.TypeOf(array3)) // [5]int

	slice = append(slice, 18) // adicionar valores ao slice
	fmt.Println(slice) // [10 11 12 13 14 15 16 17 18]

	slice2 := array2[1:3] // Funciona como um ponteiro
	fmt.Println(slice2) // [Raphael Cabral]

	// Arrays internos
	fmt.Println("-----------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity



}