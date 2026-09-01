-- +goose Up
CREATE TABLE feed_follows(
    id UUID primary key DEFAULT gen_random_uuid(),
    created_at timestamp not null DEFAULT now(),
    updated_at timestamp not null DEFAULT now(),
    user_id UUID not null references users(id) ON DELETE CASCADE,
    feed_id UUID not null references feeds(id) ON DELETE CASCADE,
    unique(user_id, feed_id)
);
-- +goose Down
DROP TABLE feed_follows;