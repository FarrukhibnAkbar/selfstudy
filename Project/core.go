package main

import(
	"encoding/json"
	"net/http"
)


func Basic(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	encoder.Encode([]string{"Osh", "Mastava", "Qozon Kabob", "Ko'k Somsa"})
}