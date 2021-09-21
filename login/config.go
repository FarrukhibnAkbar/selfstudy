package main

import(
	"fmt"
)

var(
	Host="localhost"
	User="admin"
	Password="12345"
	Database="accaunt"
	Port=5432
)

var DB_CONFIG = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%d",
	Host, User, Password, Database, Port,
)

