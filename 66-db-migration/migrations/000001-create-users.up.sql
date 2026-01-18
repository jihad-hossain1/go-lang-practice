-- +migrate Up


CREATE TABLE if not exists users (
    id SERIAL PRIMARY KEY,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    email  varchar(255) UNIQUE NOT NULL,
    password varchar(255) not null,
    is_shop_owner boolean default false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
