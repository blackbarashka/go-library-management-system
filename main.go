package main

import (
	myhttp "lib/http"
	"lib/library"
	"log"
	"net/http"
)

func main() {
	lib := library.NewLibrary()
	handler := myhttp.NewHandler(lib)
	mux := http.NewServeMux()

	mux.HandleFunc("POST /books", handler.HttpAdd)
	mux.HandleFunc("GET /books/{id}", handler.HttpGetBook)
	mux.HandleFunc("GET /books", handler.HttpFiltredGet)
	mux.HandleFunc("PATCH /books/{id}", handler.HttpMarkAsRead)
	mux.HandleFunc("DELETE /books/{id}", handler.HttpDeleteBook)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
