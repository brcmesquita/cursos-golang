package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	// // For forma simples
	for i < 10 {
		i++
		fmt.Println(i, "e contando...")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	// // For forma comum
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando o i:", j)
		time.Sleep(time.Second)
	}

	// Iterando por um array
	nomes := [4]string{"Bruno", "Raphael", "Cabral", "Mesquita"}

	n := 0
	for n < len(nomes) {
		fmt.Println(nomes[n])
		n++
	}

	// Iterando por um array, forma 2
	// Parece muito com um forEach
	// ele separa por chave (indice) e valor (nome)
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	// 0 Bruno
	// 1 Raphael
	// 2 Cabral
	// 3 Mesquita

	// Para não mostrar o índice, vocẽ usa um underline ao invés de escrever índice
	// E na hora de printar, vc só coloca o nome
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// Iterar por um slice
	nomes2 := [4]string{"Bruno", "Raphael", "Cabral", "Mesquita"}

	for _, nome := range nomes2 {
		fmt.Println("-",nome)
	}

	// Iterar por uma string
	for indice, letra := range "Raphael" {
		fmt.Println(indice, string(letra))
	}

	// iterar por um map
	usuario := map[string]string {
		"nome": "Stéphanie",
		"sobrenome": "Nunes",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// For infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}