package main

import (
	"fmt"
	"time"
)

func main() {
	// select - espera por múltiplas operações de canal
	// A instrução select bloqueia até que uma das suas operações de canal possam ser executadas, então ela executa essa operação. Se múltiplas operações puderem ser executadas, uma delas é escolhida aleatoriamente.
	canal1, canal2 := make(chan string), make(chan string)

	go func() { // goroutine 1
		for {
			time.Sleep(time.Second * 500)
			canal1 <- "Mensagem do canal 1"
		}
	}()

	go func() { // goroutine 2
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Mensagem do canal 2"
		}
	}()

	for { // loop principal
		select { // Espera por mensagens em múltiplos canais
		case mensagem1 := <-canal1: // Se receber do canal1
			fmt.Println(mensagem1)

		case mensagem2 := <-canal2: // Se receber do canal2
			fmt.Println(mensagem2)
		}
	}

}
