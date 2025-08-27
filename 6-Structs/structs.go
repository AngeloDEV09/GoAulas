package main

import "fmt"

// struct = estrutura
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logadouro string
	numero    uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Ruan"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Angelo", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Lucas"}
	fmt.Println(usuario3)
}
