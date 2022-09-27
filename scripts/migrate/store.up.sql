BEGIN;

CREATE TABLE if NOT EXISTS store.user
(
    id integer NOT NULL AUTO_INCREMENT,
    email text unique NOT NULL,
    password text  NOT NULL,
    role text NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE if NOT EXISTS store.basket
(
    id integer NOT NULL AUTO_INCREMENT,
    user_id text NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE if NOT EXISTS store.basket_product
(
    id integer NOT NULL AUTO_INCREMENT,
    product_id text NOT NULL,
    basket_id text NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE if NOT EXISTS store.product
(
    id integer NOT NULL AUTO_INCREMENT,
    name text unique NOT NULL,
    price real NOT NULL,
    img text NOT NULL,
    type_id text NOT NULL,
    brand_id text NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE if NOT EXISTS store.product_info
(
    id integer NOT NULL AUTO_INCREMENT,
    product_id text NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE if NOT EXISTS store.type
(
    id integer NOT NULL AUTO_INCREMENT,
    name text NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE if NOT EXISTS store.brand
(
    id integer NOT NULL AUTO_INCREMENT,
    name text NOT NULL,
    PRIMARY KEY(id)
);


COMMIT;