-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    title VARCHAR(200),
    description TEXT,
    url TEXT,
    published_at TIMESTAMP,
    feed_id UUID
);

-- +goose Down
DROP TABLE IF EXISTS posts;
