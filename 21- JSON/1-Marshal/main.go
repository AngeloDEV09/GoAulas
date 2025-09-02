package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Estrutura para representar um cachorro
type cachorro struct {
	Nome  string `json:"nome"`
	Idade uint   `json:"idade"`
	Raca  string `json:"raca"`
}

// Função principal
func main() {
	c := cachorro{"Rex", 10, "Pastor Alemão"}
	fmt.Println(c)

	// Convertendo para JSON
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	// Imprimindo o JSON
	fmt.Println(string(cachorroEmJSON))
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	// Usando um map para representar o cachorro 2
	c2 := map[string]string{
		"nome": "Lulu",
		"raca": "Pastor Alemão",
	}
	// Convertendo o map para JSON
	cachorroEmJSON2, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	// Imprimindo o JSON
	fmt.Println(string(cachorroEmJSON2))
	fmt.Println(bytes.NewBuffer(cachorroEmJSON2))
}
