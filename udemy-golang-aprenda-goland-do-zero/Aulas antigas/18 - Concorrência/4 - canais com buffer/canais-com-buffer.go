package main

import "fmt"

func main() {
	// Canais
	// Exemplo de erros que podem acontecer com o uso de Canais

	canal := make(chan string, 2)
	// esse 2 é para informar que eu estou esperando 2 canais

	// enviando os 2 canais
	canal <- "Olá mundo!"         // canal 1
	canal <- "Programando em Go!" // canal 2

	// colocando os canais em uma variável
	mensagem := <-canal
	mensagem2 := <-canal

	// imprimindo os dois canais
	fmt.Println(mensagem)
	fmt.Println(mensagem2)

	// caso eu insira outro valor diferente de 2, vai quebrar o programa
	// caso eu adicione mais de um canal, e só foram instanciados 2, vai quebrar

}
