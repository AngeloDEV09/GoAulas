package main

// PROGRAMA PRINCIPAL
import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() { //import "introducao-testes/enderecos"
	TipoDeEndereco := enderecos.TipoDeEndereco("Rua ABC")
	fmt.Println(TipoDeEndereco)
}
