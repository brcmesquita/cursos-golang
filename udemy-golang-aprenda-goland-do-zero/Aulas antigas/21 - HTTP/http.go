package main

import (
	"log"
	"net/http"
)

func raiz(w http.ResponseWriter, r *http.Request) { // raiz
	w.Write([]byte("Página Raiz!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func cadastro(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Crie uma conta, ou faça login!"))
}

func main() {
	// HTTP é um protocolo de comunicação - Base da comunicação Web

	// Cliente x Servidor

	// Request x Response

	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/", raiz)
	http.HandleFunc("/home", home)
	http.HandleFunc("/cadastro", cadastro)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
