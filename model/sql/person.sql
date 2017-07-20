/*
    Table: Person
*/

create table person (
    id int not null auto_increment primary key,
    name varchar(32) not null,
    surname varchar(32)
)