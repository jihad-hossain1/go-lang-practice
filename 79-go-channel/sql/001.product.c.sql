create table products (
    id bigserial primary key,
    title varchar(255) not null,
    description text,
    price double precision not null, 
    img_url text,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp
)