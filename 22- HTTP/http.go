package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Página de Usuários"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}

// HTTP É UM PROTOCOLO DE COMUNICAÇÃO ENTRE CLIENTE E SERVIDOR NA WEB
// CLIENTE FAZ UMA REQUISIÇÃO E O SERVIDOR RESPONDE
// REQUEST-RESPONSE CYCLE
// ROTAS SÃO OS CAMINHOS QUE O SERVIDOR PODE ATENDER
// URI - IDENTIFICADOR UNIFORME DE RECURSOS
// METODO - GET, POST, PUT, DELETE, PATCH (VERBOS HTTP)
