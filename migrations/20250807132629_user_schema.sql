-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

CREATE TABLE dco.users (
user_id    UUID PRIMARY KEY,
email      TEXT NOT NULL UNIQUE,
password   BYTEA  NOT NULL,
name       TEXT NOT NULL,
surname    TEXT NOT NULL,
role       TEXT NOT NULL DEFAULT 'user',
created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
is_active  BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE INDEX idx_users_email ON dco.users(email);

-- +goose StatementEnd
