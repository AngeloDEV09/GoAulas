package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Multiplexador - combina múltiplas fontes de dados em uma única

func main() {
	canal := Multiplexar(escrever("Goroutine 1"), escrever("Goroutine 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal) // Recebe do canal multiplexado
	}
}

func Multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string { // Recebe dois canais de entrada e retorna um canal de saída
	canalDeSaida := make(chan string)

	go func() { // goroutine multiplexadora
		for {
			select {
			case mensagem := <-canalDeEntrada1: // Recebe de qualquer canal de entrada
				canalDeSaida <- mensagem // Envia para o canal de saída
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string { // Generator
	canal := make(chan string) // Cria um canal

	go func() { // goroutine
		for { // Loop infinito
			canal <- fmt.Sprintf("Valor: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000))) // Dorme por um tempo aleatório
		}
	}()

	return canal
}
