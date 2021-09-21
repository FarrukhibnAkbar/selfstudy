package main

import(
	"encoding/json"
	"fmt"
)

type User struct{ 
	LastName 	string		`json: "lastName"`
	FirstName 	string		`json: "firstName"`
	Age  		int 		`json: "age"`
	IsActive	bool		`json: "isActive"`
	Books		[]string	`json: "b"`
}


var json_content = `
{
	"lastName" : "Ismoilov",
	"firstName": "Farrux",
	"age"	   : 19,
	"isActive" : true,
	"b"		   : ["Sohilsiz Dengiz"]
}
`

func main() {

 	user := User{}

	err := json.Unmarshal([]byte(json_content), &user)

	if err != nil{
		
		panic(err)
	}

	fmt.Println(user.FirstName, user)
}