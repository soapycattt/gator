-- +goose Up
ALTER TABLE feeds
DROP COLUMN IF EXISTS rss_url;

-- +goose Down
ALTER TABLE feeds
ADD rss_url VARCHAR(200);
