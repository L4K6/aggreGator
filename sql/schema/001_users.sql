-- +goose Up
Create Table users(
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text unique not null
);
-- +goose Down
Drop Table users;