-- +goose Up
CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    expires_at TIMESTAMP NOT NULL,
    refresh_token TEXT UNIQUE NOT NULL
);


-- +goose Down
DROP TABLE tokens;