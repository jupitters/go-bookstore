package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jupitters/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
