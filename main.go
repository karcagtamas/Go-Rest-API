package main

import (
	"log"

	"github.com/karcatamas/restapi/handlers"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()
	handlers.Load()

	// Route Handlers / Endpoint
	r.HandleFunc("/api/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/api/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", handlers.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
