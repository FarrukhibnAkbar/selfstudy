package main

import(
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)


type Page struct{
	Title, Content string
}



func Home(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("template.html"))

	d := Page{
		Title: "Hello World",
		Content:`
			Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
			tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
			quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
			consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
			cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
			proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

		`,
	}

	t.Execute(w, d)
}


func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", Home)

	http.ListenAndServe(":8080", r)

 }
