package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ContrerasJoel/back_honeyhot/internal/product"
	"github.com/gorilla/mux"
)

func main() {
	url := "https://api.ipify.org?format=json"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	conten, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	log.Println(string(conten))

	port := os.Getenv("PORT")
	if port == "" {
		port = "7071"
	}

	r := mux.NewRouter()

	product.NewHandler().Expose(r)

	srv := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: r,
	}

	log.Println("Escuchando en el puerto " + srv.Addr)
	srv.ListenAndServe()

}
