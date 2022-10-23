BEGIN;

CREATE TABLE if NOT EXISTS store.product
(
    id          SERIAL PRIMARY KEY NOT NULL,
    sku         text               NOT NULL,
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

CREATE TABLE if NOT EXISTS store.user
(
    id          SERIAL PRIMARY KEY  NOT NULL,
    email       varchar(320) UNIQUE NOT NULL,
    username    varchar(256) UNIQUE NOT NULL,
    password    varchar(128)        NOT NULL,
    role        varchar             NOT NULL,
    created_at  timestamp           NOT NULL,
    modified_at timestamp           NOT NULL
);

CREATE TABLE if NOT EXISTS store.salt
(
    id      SERIAL PRIMARY KEY NOT NULL,
    user_id integer UNIQUE     NOT NULL,
    salt    text               NOT NULL
);

CREATE TABLE if NOT EXISTS store.persistent_logins
(
    email     varchar(320) NOT NULL,
    series    text PRIMARY KEY,
    token     text         NOT NULL,
    last_used timestamp    NOT NULL
);

CREATE TABLE if NOT EXISTS store.user_address
(
    id            SERIAL PRIMARY KEY NOT NULL,
    user_id       integer            NOT NULL,
    address_line1 text               NOT NULL,
    address_line2 text,
    city          text               NOT NULL,
    postal_code   text               NOT NULL,
    telephone     text,
    mobile        text
);

CREATE TABLE if NOT EXISTS store.order
(
    id              SERIAL PRIMARY KEY NOT NULL,
    user_id         integer            NOT NULL,
    payment_id      integer,
    total           real               NOT NULL,
    order_date      text,
    status          text               NOT NULL,
    delivery_fee    real               NOT NULL,
    tracking_number text,
    created_at      timestamp          NOT NULL,
    modified_at     timestamp          NOT NULL
);

CREATE TABLE if NOT EXISTS store.order_items
(
    id          SERIAL PRIMARY KEY NOT NULL,
    order_id    integer            NOT NULL,
    product_id  integer,
    quantity    smallint           NOT NULL,
    created_at  timestamp          NOT NULL,
    modified_at timestamp          NOT NULL
);

CREATE TABLE if NOT EXISTS store.cart
(
    id      SERIAL PRIMARY KEY NOT NULL,
    user_id integer            NOT NULL
);

CREATE TABLE if NOT EXISTS store.cart_items
(
    id          SERIAL PRIMARY KEY NOT NULL,
    cart_id     integer            NOT NULL,
    product_id  integer            NOT NULL,
    quantity    smallint           NOT NULL,
    created_at  timestamp          NOT NULL,
    modified_at timestamp          NOT NULL
);


COMMIT;