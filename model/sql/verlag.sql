/*
    Table: Verlag
*/

create table verlag(
    id int not null auto_increment primary_key,
    name varchar(32) not null,
    zusatz varchar(32),
    strasse varchar(32),
    plz varchar(10),
    ort varchar(32)
);