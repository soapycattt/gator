// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Name      sql.NullString
	Url       sql.NullString
	UserID    uuid.NullUUID
	RssUrl    sql.NullString
}

type FeedFollow struct {
	ID        uuid.UUID
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	UserID    uuid.NullUUID
	FeedID    uuid.NullUUID
}

type User struct {
	ID        uuid.UUID
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Name      sql.NullString
}