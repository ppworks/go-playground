-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE posts (
    id text NOT NULL PRIMARY KEY,
    incremental_key SERIAL UNIQUE,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    author text NOT NULL,
    content text NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE posts;
