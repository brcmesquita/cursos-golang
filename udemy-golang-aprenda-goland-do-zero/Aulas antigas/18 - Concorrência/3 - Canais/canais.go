package main

import (
	"fmt"
	"time"
)

func main() {
	// CANAIS

	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// // maneira tradicional
	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	// maneira refatorada
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	mensagem := "Fim do programa"
	fmt.Println(mensagem)
}

func escrever(texto string, canal chan string) {
	i := 0
	for i < 5 {
		canal <- texto
		time.Sleep(time.Second)
		i++
	}

	close(canal)
}
