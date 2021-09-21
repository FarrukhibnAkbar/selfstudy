package main

import(
	"net/http"
)

func main(){
	
	http.HandleFunc("/", Basic)

	http.ListenAndServe(":8008", nil)
}