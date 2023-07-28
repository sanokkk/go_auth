-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(128) NOT NULL UNIQUE,
    nick_name VARCHAR(64) NOT NULL UNIQUE,
    age SMALLINT NOT NULL,
    password_hash TEXT NOT NULL
);

-- +goose Down
CREATE TABLE users;