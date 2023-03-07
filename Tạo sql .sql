create table users
(
    id        int auto_increment,
    name      varchar(32) null,
    age       int         null,
    home_town varchar(32) null,
    password  varchar(32) null,
    username  varchar(32) null,
    constraint new_pk
        primary key (id)
);