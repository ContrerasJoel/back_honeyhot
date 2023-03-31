package main

import (
	"log"
	"net/http"

	"github.com/ContrerasJoel/back_honeyhot/internal/product"
	"github.com/gorilla/mux"
)

func main() {
	port := "7071"

	r := mux.NewRouter()

	product.NewHandler().Expose(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Println("Escuchando en el puerto http://localhost:" + port)
	srv.ListenAndServe()

}
