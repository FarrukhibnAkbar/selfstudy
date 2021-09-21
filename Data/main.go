package main

import(
	"fmt"
	"net/http"
	// "html/template"
	"github.com/gorilla/mux"
)


func Home(w http.ResponseWriter, r *http.Request){
	
}


func main(){

	fmt.Println(GetUsers())

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8080",r)

}