-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    name VARCHAR(50) UNIQUE
);

-- +goose Down
DROP TABLE IF EXISTS users;
