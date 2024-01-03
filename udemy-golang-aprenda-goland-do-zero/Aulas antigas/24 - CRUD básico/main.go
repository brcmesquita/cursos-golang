package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD = Create, Read, Update, Delete

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)          // Create
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)         // Read
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)     // Read
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)  // Alter
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete) // Delete user

	fmt.Println("Escutando a porta 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))

}
