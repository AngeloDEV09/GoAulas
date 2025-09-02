package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Idade uint   `json:"idade"`
	Raca  string `json:"raca"`
}

// Unmarshal - Deserialização de JSON
func main() {
	cachorroEmJSON := `{"nome":"Rex","idade":5,"raca":"Pastor Alemão"}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	// Outro exemplo
	// Deserializando em um map
	cachorroEmJSON2 := `{"nome":"Bob","raca":"Labrador"}`

	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorroEmJSON2), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)

}
