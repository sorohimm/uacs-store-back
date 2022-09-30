BEGIN;

CREATE TABLE if NOT EXISTS store.product
(
    id SERIAL PRIMARY KEY NOT NULL,
    name text unique NOT NULL,
    price real NOT NULL,
    img text,
    type_id text NOT NULL,
    brand_id text NOT NULL
);

CREATE TABLE if NOT EXISTS store.product_stock
(
    id SERIAL PRIMARY KEY NOT NULL,
    product_id integer NOT NULL,
    in_stock boolean,
    stock_status text,
    quantity_stock integer NOT NULL
);

CREATE TABLE if NOT EXISTS store.product_info
(
    id SERIAL PRIMARY KEY NOT NULL,
    product_id integer NOT NULL,
    title text NOT NULL,
    description text NOT NULL
);

CREATE TABLE if NOT EXISTS store.category
(
    id SERIAL PRIMARY KEY NOT NULL,
    name text NOT NULL
);

CREATE TABLE if NOT EXISTS store.brand
(
    id SERIAL PRIMARY KEY NOT NULL,
    name text NOT NULL
);


COMMIT;