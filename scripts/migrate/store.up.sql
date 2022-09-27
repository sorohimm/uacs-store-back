BEGIN;

create table if not exists store.user
(
    id text not null,
    email text not null,
    password text  not null,
    role text not null
);

create table if not exists store.basket
(
    id text not null,
    user_id text not null
);

create table if not exists store.basket_product
(
    id text not null,
    product_id text not null,
    basket_id text not null
);

create table if not exists store.product
(
    id text not null,
    name text not null,
    price text not null,
    img text not null,
    type_id text not null,
    brand_id text not null
);

create table if not exists store.product_info
(
    id text not null,
    product_id text not null,
    title text not null,
    description text not null,
);

create table if not exists store.type
(
    id text not null,
    name text not null
);

create table if not exists store.brand
(
    id text not null,
    name text not null
);


COMMIT;