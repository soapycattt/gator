-- name: CreateFeeds :one
INSERT INTO feeds (
    id,
    created_at,
    updated_at,
    name,
    url,
    user_id
)

VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
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
SELECT * FROM feeds WHERE url = $1;


-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetNextFeedToFetch :many
SELECT *
FROM feeds 
ORDER BY last_fetched_at DESC NULLS FIRST;
