package main

import(
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)

var(
	Host ="localhost"
	Username="admin"
	Password=12345
	Database="mydb"
	Port=5432
)

const QUERY = "select user_id, name from users"

type User struct{
	UserId   int
	Fullname string
}

func GetUsers() []User{

	c := fmt.Sprintf(
		"host=%s user=%s password=%d dbname=%s port=%d",
		Host, Username, Password, Database, Port,
	)

	db, err := sql.Open("postgres", c)

	if err != nil{
		panic(err)
	}

	defer db.Close()

	var user User

	// err	= db.Query(QUERY).Scan(&user.UserId, &user.Fullname)
	rows, _ := db.Query(QUERY)

	defer rows.Close()

	var users []User

	for rows.Next() {

		rows.Scan(&user.UserId, &user.Fullname)

		users = append(users, user)
	}
	return users
}