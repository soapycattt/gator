-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    user_id UUID,
    feed_id UUID,
    UNIQUE(user_id, feed_id),

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (feed_id) REFERENCES feeds (id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE IF EXISTS feed_follows ;
