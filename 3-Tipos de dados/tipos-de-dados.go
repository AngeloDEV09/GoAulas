package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64 tipos de inteiro
	var numero int16 = 100
	fmt.Println(numero)
	var numero2 int = 1000
	fmt.Println(numero2)
	numero3 := -100000000000
	fmt.Println(numero3)

	//uint = sem sinal
	var numero4 uint = 1000000
	fmt.Println(numero4)

	// alias= apelido
	// int32 = rune
	var numero5 rune = 12344
	fmt.Println(numero5)
	// uint8 = byte
	var numero6 byte = 123
	fmt.Println(numero6)

	// REAIS = float32 float64 aceita ponto
	var numeroReal1 float32 = 123.43
	fmt.Println(numeroReal1)

	numeroReal2 := 123443.43
	fmt.Println(numeroReal2)

	// str string = cadeia de caracter

	var str string = "texto"
	fmt.Println(str)

	str2 := "Ruan"
	fmt.Println(str2)

	// Tabela ask
	char := 'B'
	fmt.Println(char)

	// FIM STRINGS

	var texto int
	fmt.Println(texto)

	// booleano e erro

	var booleano1 bool
	fmt.Println(booleano1)

	var booleano2 bool = true
	fmt.Println(booleano2)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}
