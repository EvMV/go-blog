-- +goose Up
-- SQL в этом блоке будет применен при применении миграции
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

-- +goose Down
-- SQL в этом блоке будет применен при откате миграции
DROP TABLE users;
