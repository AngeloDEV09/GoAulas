package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRENCIA != PARALELISMO
	// Concorrência é quando temos múltiplas tarefas que podem ser iniciadas, execut
	// PARALELISMO é quando essas tarefas são realmente executadas ao mesmo tempo.
	go escrever("Ola Mundo") // goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {

	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
