BEGIN;

create table if not exists store.user
(
    id integer not null AUTO_INCREMENT,
    email text unique not null,
    password text  not null,
    role text not null,
    primary key(id)
);

create table if not exists store.basket
(
    id integer not null AUTO_INCREMENT,
    user_id text not null,
    primary key(id)
);

create table if not exists store.basket_product
(
    id integer not null AUTO_INCREMENT,
    product_id text not null,
    basket_id text not null,
    primary key(id)
);

create table if not exists store.product
(
    id integer not null AUTO_INCREMENT,
    name text unique not null,
    price real not null,
    img text not null,
    type_id text not null,
    brand_id text not null,
    primary key(id)
);

create table if not exists store.product_info
(
    id integer not null AUTO_INCREMENT,
    product_id text not null,
    title text not null,
    description text not null,
    primary key(id)
);

create table if not exists store.type
(
    id integer not null AUTO_INCREMENT,
    name text not null,
    primary key(id)
);

create table if not exists store.brand
(
    id integer not null AUTO_INCREMENT,
    name text not null,
    primary key(id)
);


COMMIT;