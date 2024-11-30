// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
    id,
    created_at,
    updated_at,
    published_at,
    title,
    description,
    url,
    feed_id
)

VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING id, created_at, updated_at, title, description, url, published_at, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	PublishedAt sql.NullTime
	Title       sql.NullString
	Description sql.NullString
	Url         sql.NullString
	FeedID      uuid.NullUUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.PublishedAt,
		arg.Title,
		arg.Description,
		arg.Url,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Description,
		&i.Url,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const getPostsForUser = `-- name: GetPostsForUser :many
SELECT posts.id, posts.created_at, posts.updated_at, posts.title, posts.description, posts.url, posts.published_at, posts.feed_id
FROM posts 
JOIN feed_follows
  ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY published_at DESC
LIMIT $2
`

type GetPostsForUserParams struct {
	UserID uuid.NullUUID
	Limit  int32
}

func (q *Queries) GetPostsForUser(ctx context.Context, arg GetPostsForUserParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsForUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Description,
			&i.Url,
			&i.PublishedAt,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}