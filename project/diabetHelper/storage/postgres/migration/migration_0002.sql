create table if not exists users(
    id              int primary key,
    username        varchar(40),
    first_name      varchar(40),
    last_name       varchar(40),
    language_code   varchar(5)
);