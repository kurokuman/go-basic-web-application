package main

import (
	"handlers"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	WEBSERVERPORT = ":8080"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/signup", handlers.SignupHandler).Methods("GET", "POST")

	http.Handle("/", r)
	http.ListenAndServe(WEBSERVERPORT, nil)
}
