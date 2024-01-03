package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// WAITGROUP

	var waitGroup sync.WaitGroup // declaro uma variável do tipo WaitGroup

	waitGroup.Add(2) // informo ao waitGroup que possuo 2

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	i := 0
	for i < 5 {
		fmt.Println(texto)
		time.Sleep(time.Second)
		i++
	}
}
