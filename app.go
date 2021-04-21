package main

import (
	"net/http"

	"gawds/src/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", routes.CreateUser).Methods("POST")
	r.HandleFunc("/user/{key}", routes.GetUser).Methods("GET")
	http.Handle("/", r)
}
