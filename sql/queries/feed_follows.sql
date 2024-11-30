-- name: CreateFeedFollow :one
INSERT INTO feed_follows  (
    id,
    created_at,
    updated_at,
    user_id,
    feed_id
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetFeedFollowsForUser :many
SELECT 
  feeds.name as feed
FROM feed_follows
LEFT JOIN users
  ON feed_follows.user_id = users.id
LEFT JOIN feeds
  ON feed_follows.feed_id = feeds.id
WHERE users.name = $1;


-- name: DeleteFollow :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;
