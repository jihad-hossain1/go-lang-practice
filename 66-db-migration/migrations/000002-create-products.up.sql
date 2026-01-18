-- +migrate Up

CREATE TABLE if not exists products (
    id SERIAL PRIMARY KEY,
    title varchar(255) not null,
    description text,
    price decimal(10,2) not null,
    img_url varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);