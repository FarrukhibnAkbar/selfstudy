package main

import(
	"net/http"
	"encoding/json"
	// "fmt"
)


func Numbers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	

	encoder := json.NewEncoder(w)

	encoder.Encode([]int{ 10 ,20, 30, 40, 50 })
}

func Strings(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	encoder.Encode([]string{"Apple", "Facebook", "Google", "Microsoft"})

}



func main() {

	http.HandleFunc("/", Numbers)

	http.ListenAndServe(":8000", nil)
}