package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)      // Cria um canal de comunicação entre goroutines
	go escrever("Ola Mundo", canal) // goroutine

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <-canal // Recebe o texto do canal
		if !aberto {
			break // Sai do loop se o canal estiver fechado
		}
		fmt.Println(mensagem)
	}

	for mensagem := range canal { // Outra forma de ler do canal até ele ser fechado
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Envia o texto para o canal
		time.Sleep(time.Second)
	}

	close(canal) // Fecha o canal (opcional, mas uma boa prática)

}

// deadlock - todos estão esperando algo que nunca vai acontecer
