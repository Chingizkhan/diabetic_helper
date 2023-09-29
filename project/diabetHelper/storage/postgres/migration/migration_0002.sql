create table if not exists users(
    id              varchar(16) primary key,
    username        varchar(40),
    first_name      varchar(40),
    last_name       varchar(40),
    language_code   varchar(5)
);