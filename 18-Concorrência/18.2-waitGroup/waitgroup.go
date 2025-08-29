package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// CONCORRENCIA != PARALELISMO
	// Concorrência é quando temos múltiplas tarefas que podem ser iniciadas, execut
	// PARALELISMO é quando essas tarefas são realmente executadas ao mesmo tempo.
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // Adiciona o número de goroutines que vamos esperar

	go func() {
		escrever("Ola Mundo")
		waitGroup.Done() // Decrementa o contador do waitGroup -1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // Decrementa o contador do waitGroup -1
	}()

	waitGroup.Wait() // Espera todas as goroutines terminarem -2 = 0
}

func escrever(texto string) {

	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
