package main

import (
	"net/http"

	"gawds/src/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.CreateUser)
	r.HandleFunc("/products", routes.GetUser)
	http.Handle("/", r)
}
