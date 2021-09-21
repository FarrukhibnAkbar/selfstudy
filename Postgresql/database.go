package main

var(
	Host ="localhost"
	User="admin"
	Password=12345
	Database="web"
	Port=5432
)

type User struct{
	UserId   int
	Fullname string
}