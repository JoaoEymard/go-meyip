package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./src/router"
)

func main() {
	r := router.Configuration(mux.NewRouter())

	fmt.Println("Servidor iniciado")
	log.Fatal(http.ListenAndServe(":5000", r))
}
