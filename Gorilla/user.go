package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"sort"
	"encoding/json"
	"errors"
	"fmt"
)


type User struct{

	Name string
	Age  int
}


var users []User =[]User{
	User{ "Farrux" , 19 },
	User{ "Anvar"  , 20 },
	User{ "Suhrot" , 26 },
	User{ "Shahzod", 18 },
}


func GetUsers(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)
	v := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if v["y"] == "asc"{

		sort.Slice(users, func(i, j int) bool {
			return users[i].Name < users[j].Name
		})
			
		encoder.Encode(users)
		
	} else if v["y"] == "desc" {

		sort.Slice(users, func(i ,j int) bool {
			return users[i].Name > users[j].Name
		})

		encoder.Encode(users)

	} else {

		fmt.Println(errors.New("http: named cookie not present"))
	}
}


func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("User Profile"))
}


