-- name: CreatePost :one

INSERT INTO posts(
    id,
    created_at,
    updated_at,
    title,
    url,
    description,
    published_at,
    feed_id
) VALUES(
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: GetPostsForUser :many

SELECT p.*
FROM posts AS p
JOIN feed_follows AS ff ON ff.feed_id = p.feed_id
JOIN users AS u ON u.id = ff.user_id
WHERE u.name = $1
ORDER BY published_at DESC LIMIT $2;

