CREATE TABLE IF NOT EXISTS user(
    id int primary key,
    email varchar(20) not null,
    password varchar(30) not null,
    name varchar(20) not null
);

CREATE TABLE IF NOT EXISTS blog(
    id serial primary key ,
    title varchar(50) not null,
    content text,
    created_at datetime default now(),
    updated_at datetime default now(),
    user int not null,
    foreign key (user) references user(id)
);

