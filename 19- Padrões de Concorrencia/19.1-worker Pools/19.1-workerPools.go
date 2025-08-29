package main

func main() {
	tarefas := make(chan int, 45)    // Canal para enviar tarefas (posições da sequência de Fibonacci)
	resultados := make(chan int, 45) // Canal para receber resultados

	go worker(tarefas, resultados) // Inicia um trabalhador
	go worker(tarefas, resultados) // Inicia outro trabalhador
	go worker(tarefas, resultados) // Inicia mais um trabalhador

	for i := 0; i < 45; i++ {
		tarefas <- i // Envia tarefas para o canal
	}
	close(tarefas) // Fecha o canal de tarefas após enviar todas as tarefas

	for i := 0; i < 45; i++ {
		resultado := <-resultados // Recebe resultados do canal
		println(resultado)        // Imprime o resultado
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for posicao := range tarefas { // Recebe tarefas do canal de tarefas
		resultado := fibonacci(posicao) // Processa a tarefa (calcula Fibonacci)
		resultados <- resultado         // Envia o resultado para o canal de resultados
	}
}

// Worker Pools é um padrão de concorrência onde um número fixo de goroutines (trabalhadores) são criadas para processar tarefas de forma concorrente. Isso é útil para limitar o número de goroutines em execução ao mesmo tempo, evitando sobrecarga do sistema e melhorando a eficiência do processamento.
// O padrão Worker Pools é frequentemente usado em cenários onde há muitas tarefas a serem processadas, mas o número de recursos do sistema (como CPU e memória) é limitado. Ao usar um pool de trabalhadores, podemos controlar o número de tarefas que são processadas simultaneamente, garantindo que o sistema não fique sobrecarregado.

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
