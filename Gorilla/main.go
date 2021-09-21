package main

import(
	"net/http"
	"github.com/gorilla/mux"
)


func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", GetUsers).
	Methods("GET").
	Queries("x", "{y}")

	router.HandleFunc("/{username}", GetUser).Methods("GET")

	http.ListenAndServe(":8080", router) 
}