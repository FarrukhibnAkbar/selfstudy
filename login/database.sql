create table users(
	user_id serial not null primary key,
	firstname varchar(32) not null,
	lastname  varchar(32) not null,
	age int not null,
	create_at timestamp with time zone default current_timestamp
);


insert into users(firstname, lastname, age) values('Farrux', 'Ismoilov', 19);