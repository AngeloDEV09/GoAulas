package main

import "fmt"

func main() {
	canal := make(chan string, 2) // Canal com buffer de 2 posições
	canal <- "Mensagem 1"
	canal <- "Mensagem 2"
	// canal <- "Mensagem 3" // Isso causaria um deadlock, pois o buffer está cheio

	mensagem1 := <-canal
	fmt.Println(mensagem1)
	mensagem2 := <-canal
	fmt.Println(mensagem2)
	// mensagem3 := <-canal // Isso causaria um deadlock, pois não há mais mensagens no canal
}
