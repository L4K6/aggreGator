-- +goose Up
CREATE TABLE feeds(
    id UUID primary key DEFAULT gen_random_uuid(),
    created_at timestamp not null DEFAULT now(),
    updated_at timestamp not null DEFAULT now(),
    name text not null,
    url text unique not null,
    user_id UUID not null references users(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE feeds;