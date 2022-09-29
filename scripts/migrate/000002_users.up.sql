BEGIN;

CREATE TABLE if NOT EXISTS users.user
(
    id SERIAL PRIMARY KEY NOT NULL,
    email text unique NOT NULL,
    password text  NOT NULL,
    role text NOT NULL
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