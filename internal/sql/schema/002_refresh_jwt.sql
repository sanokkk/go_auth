-- +goose Up
CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    jwt_token TEXT UNIQUE NOT NULL,
    refresh_token TEXT UNIQUE NOT NULL
);


-- +goose Down
DROP TABLE tokens;