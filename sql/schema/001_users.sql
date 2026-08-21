-- +goose Up
Create Table users(
    id int primary key,
    created_at timestamp not null,
    updated_at timestamp null,
    name text unique not null
);

-- +goose Down
Drop Table users;