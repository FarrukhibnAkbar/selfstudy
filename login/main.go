package main

import(
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/registration", GetRegistration).Methods("GET")
	r.HandleFunc("/registration", PostRegistration).Methods("POST")
	r.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request){
		t := template.Must(template.ParseFiles("main.html"))

		t.Execute(w, ,GetAccaunt())
	})

	http.ListenAndServe(":8080", r)
}