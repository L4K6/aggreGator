-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    insert into feed_follows (user_id, feed_id)
    VALUES($1, $2)
    RETURNING *
)
Select inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
from inserted_feed_follow
    inner join users on inserted_feed_follow.user_id = users.id
    inner join feeds on inserted_feed_follow.feed_id = feeds.id;