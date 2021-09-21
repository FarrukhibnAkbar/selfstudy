create table users(
	user_id serial not null primary key,
	username varchar(32) not null,
	firstname varchar(32) not null,
	lastname varchar(32) not null
	
);

insert into users(username, firstname, lastname) values
('husniddin', 'Husniddin', 'Ikramov'),
('farrux', 'Farrux', 'Ismoilov')
;

insert into users(username, firstname, lastname) values
('jasur', 'Jasur', 'Shamsiddinov'),
('asadbek', 'Asadbek', 'Ergashov')
;