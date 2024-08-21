create table members (
    id serial primary key,
    name varchar(255),
    age int,
    sex varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now()
);
