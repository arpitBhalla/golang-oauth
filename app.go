package main

import (
	"log"
	"net/http"

	"gawds/src/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", routes.CreateUser).Methods("POST")
	r.HandleFunc("/user/{key}", routes.GetUser).Methods("GET")

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Server Running"))
	})

	log.Println("Server Stated")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("Error Occured", err)
	}
}
