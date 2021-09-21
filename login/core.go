package main

import(
	"database/sql"
	_"github.com/lib/pq"
)

type Accaunts struct{
	Id 		  int
	Firstname string
	Lastname  string
	Age 	  int
	NowTime   string
}


type AccauntsBody struct{
	Firstname string
	Lastname  string
	Age       int
}


const GET_QUERY = `
	select
		user_id,
		firstname,
		lastname,
		age,
		create_at
	from users
` 


const POST_QUERY =`
	insert into users(firstname, lastname, age)values($1, $2, $3)
	returning
		user_id,
		firstname,
		lastname,
		age,
		create_at

`


func GetAccaunt() []Accaunts{

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()
	
	if err != nil{
		panic(err)
	}

	rows, err := db.Query(GET_QUERY)

	defer rows.Close()

	if err != nil {
		panic(err)
	}
	
	var accaunts []Accaunts

	for rows.Next(){
		
		var accaunt Accaunts = Accaunts{}

		rows.Scan(
			&accaunt.Id, 
			&accaunt.Firstname, 
			&accaunt.Lastname, 
			&accaunt.Age, 
			&accaunt.NowTime)

		accaunts = append(accaunts, accaunt)
	}

	return accaunts
}



func PostAccaunt(body AccauntsBody) Accaunts {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		panic(err)
	} 

	var newAccaunts Accaunts

	err = db.QueryRow(POST_QUERY, body.Firstname, body.Lastname, body.Age).Scan(&newAccaunts.Id, &newAccaunts.Firstname, &newAccaunts.Lastname, &newAccaunts.Age, &newAccaunts.NowTime) 

	if err != nil{
		panic(err)
	}

	return newAccaunts
}	








