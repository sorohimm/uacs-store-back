BEGIN;

CREATE TABLE if NOT EXISTS store.product
(
    id          SERIAL PRIMARY KEY NOT NULL,
    name        text unique        NOT NULL,
    price       real               NOT NULL,
    img         text,
    type_id     text,
    brand_id    text,
    created_at  timestamp          NOT NULL,
    modified_at timestamp          NOT NULL,
    deleted_at  timestamp
);

CREATE TABLE if NOT EXISTS store.product_stock
(
    id             SERIAL PRIMARY KEY NOT NULL,
    product_id     integer            NOT NULL,
    in_stock       boolean,
    stock_status   text,
    quantity_stock integer            NOT NULL
);

CREATE TABLE if NOT EXISTS store.product_info
(
    id          SERIAL PRIMARY KEY NOT NULL,
    product_id  integer            NOT NULL,
    title       text               NOT NULL,
    description text               NOT NULL
);

CREATE TABLE if NOT EXISTS store.category
(
    id          SERIAL PRIMARY KEY NOT NULL,
    name        text               NOT NULL,
    created_at  timestamp          NOT NULL,
    modified_at timestamp          NOT NULL,
    deleted_at  timestamp
);

CREATE TABLE if NOT EXISTS store.brand
(
    id          SERIAL PRIMARY KEY NOT NULL,
    name        text               NOT NULL,
    created_at  timestamp          NOT NULL,
    modified_at timestamp          NOT NULL,
    deleted_at  timestamp
);

CREATE TABLE if NOT EXISTS store.orders
(
    id              SERIAL PRIMARY KEY NOT NULL,
    user_id         text               NOT NULL,
    total           real               NOT NULL,
    order_date      text,
    status          text               NOT NULL,
    delivery_fee    real               NOT NULL,
    tracking_number text
);


COMMIT;