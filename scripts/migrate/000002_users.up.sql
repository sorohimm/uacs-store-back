BEGIN;

CREATE TABLE if NOT EXISTS users.user
(
    id SERIAL PRIMARY KEY NOT NULL,
    email varchar(64) UNIQUE NOT NULL,
    password text NOT NULL,
    role text NOT NULL
);

CREATE TABLE if NOT EXISTS users.salt
(
    id SERIAL PRIMARY KEY NOT NULL,
    user_id integer UNIQUE NOT NULL,
    salt text NOT NULL,
);

CREATE TABLE if NOT EXISTS users.persistent_logins (
    email varchar(64) NOT NULL,
    series varchar(64) PRIMARY KEY,
    token varchar(64) NOT NULL,
    last_used timestamp NOT NULL
);

CREATE TABLE if NOT EXISTS users.basket
(
    id SERIAL PRIMARY KEY NOT NULL,
    user_id text NOT NULL
);

CREATE TABLE if NOT EXISTS users.basket_product
(
    id SERIAL PRIMARY KEY NOT NULL,
    product_id text NOT NULL,
    basket_id text NOT NULL
);

COMMIT;