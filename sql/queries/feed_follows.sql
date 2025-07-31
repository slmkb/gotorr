-- name: CreateFeedFollow :one
WITH inserted AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES(
        $1,
        $2,
        $3,
        $4,
        $5
    ) RETURNING * 
)
SELECT 
    inserted.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted
INNER JOIN users ON inserted.user_id = users.id
INNER JOIN feeds ON inserted.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
WITH user_cte AS (
    SELECT id AS user_id FROM users WHERE users.name = $1
)
SELECT feeds.name
FROM feed_follows
INNER JOIN user_cte ON feed_follows.user_id = user_cte.user_id
INNER JOIN feeds ON feeds.id = feed_follows.feed_id;
