package main

import (
	"fmt"
	"time"
)

// Generator é um padrão de concorrência onde uma função gera uma série de valores e os envia através de um canal para serem consumidos por outra parte do programa. Isso é útil para criar pipelines de dados ou para produzir valores sob demanda.

// O padrão Generator é frequentemente usado em cenários onde os dados são produzidos de forma assíncrona ou quando a produção de dados é cara em termos de recursos. Ao usar um generator, podemos controlar o ritmo da produção de dados e evitar sobrecarregar o consumidor.
func main() {
	canal := escrever("Ola Mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal) // Recebe o valor do canal
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
