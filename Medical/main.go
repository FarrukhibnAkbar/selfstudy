package main

import(
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)


func Home(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("front.html"))

	t.Execute(w,r)
}

func Twopage(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("second.html"))

	t.Execute(w,r)
}


func main(){

	r := mux.NewRouter()

	r.HandleFunc("/", Home)
	r.HandleFunc("/twopage", Twopage)

	http.ListenAndServe(":8080", r)
}
