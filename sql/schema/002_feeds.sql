-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    name VARCHAR(100),
    url VARCHAR(200),
    user_id UUID,

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE IF EXISTS feeds;
