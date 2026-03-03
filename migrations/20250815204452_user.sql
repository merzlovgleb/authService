-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS dco;

CREATE TABLE IF NOT EXISTS dco.users (
                           user_id    UUID PRIMARY KEY,
                           email      TEXT NOT NULL UNIQUE,
                           password   BYTEA NOT NULL,
                           name       TEXT NOT NULL,
                           surname    TEXT NOT NULL,
                           role       TEXT NOT NULL DEFAULT 'user',
                           created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           is_active  BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE INDEX IF NOT EXISTS idx_users_email ON dco.users(email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dco.users;
DROP SCHEMA IF EXISTS dco CASCADE;
-- +goose StatementEnd