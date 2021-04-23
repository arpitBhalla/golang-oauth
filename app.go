package main

import (
	"log"
	"net/http"

	"gawds/src/middleware"
	"gawds/src/routes"
	"gawds/src/utils"

	"github.com/gorilla/mux"
)

func main() {
	utils.RedisInit()
	r := mux.NewRouter()
	r.Use(middleware.Middleware)
	r.HandleFunc("/login", routes.Login).Methods("POST")
	r.HandleFunc("/logout", routes.Logout).Methods("POST")
	r.HandleFunc("/register", routes.Register).Methods("POST")
	r.HandleFunc("/profile", routes.Profile).Methods("GET")
	r.HandleFunc("/refresh", routes.RefreshToken).Methods("POST")
	r.HandleFunc("/all", routes.GetAll).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Route Not Found"))
	})
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	})
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server Running"))
	})
	log.Println("Server Stated")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("Error Occured", err)
	}
}
