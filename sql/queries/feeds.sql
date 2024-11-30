-- name: CreateFeeds :one
INSERT INTO feeds (
    id,
    created_at,
    updated_at,
    name,
    url,
    user_id,
    rss_url
)

VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;


-- name: ListFeeds :many
SELECT 
  feeds.name as feed,
  feeds.url,
  users.name as user
FROM feeds 
LEFT JOIN users
  ON feeds.user_id = users.id;



-- name: DeleteAllFeeds :exec
DELETE FROM feeds;


-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE rss_url = $1;
