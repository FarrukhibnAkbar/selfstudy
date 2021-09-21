package main

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
)


func GetRegistration(w http.ResponseWriter, r *http.Request){
	
	e := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")

	e.Encode(GetAccaunt())
}


func PostRegistration(w http.ResponseWriter, r *http.Request){

	e := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")

	var newAccaunts AccauntsBody

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &newAccaunts)

	e.Encode(PostAccaunt(newAccaunts))
}