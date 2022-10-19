BEGIN;

CREATE TABLE if NOT EXISTS users.user
(
    id SERIAL PRIMARY KEY NOT NULL,
    email text UNIQUE NOT NULL,
    username text UNIQUE NOT NULL,
    password text NOT NULL,
    role text NOT NULL
);

CREATE TABLE if NOT EXISTS users.salt
(
    id SERIAL PRIMARY KEY NOT NULL,
    user_id integer UNIQUE NOT NULL,
    salt text NOT NULL
);

CREATE TABLE if NOT EXISTS users.persistent_logins
(
    email text NOT NULL,
    series text PRIMARY KEY,
    token text NOT NULL,
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