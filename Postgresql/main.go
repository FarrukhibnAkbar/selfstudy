package main

import(
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"github.com/gorilla/mux"
)

func main(){

	c := fmt,Sprintf(
		"host=%s user=%s password=%d dbname=%s port=%d",
		Host, User, Password, Database, Port,
	)

	db, err := sql,Open("postgres", c)

	if err != nil{
		panic(err)
	}

	defer db.Close()

	fmt.Println("ok")
}