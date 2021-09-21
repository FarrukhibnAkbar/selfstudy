package main

import(
	"fmt"
	"encoding/json"
)

type Comment struct{
	UserId  	int
	Id 			int
	Title		string
	Completed	bool
}


func main() {

	comment := Comment{
		1,
		1,
		"delectus aut autem",
		false,
	}

	bytes,err := json.Marshal(comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}