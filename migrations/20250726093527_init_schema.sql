-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE SCHEMA if not exists dco;

CREATE TABLE if not exists dco.segments
(
    group_id uuid    NOT NULL,
    id       uuid    NOT NULL,
    title    text    not null,
    p        integer not null,
    PRIMARY KEY (id)
);

CREATE TABLE if not exists dco.clients_segments
(
    client_id  uuid NOT NULL,
    segment_id uuid NOT NULL,
    PRIMARY KEY (client_id, segment_id)
);

CREATE TABLE if not exists dco.groups
(
    id    uuid NOT NULL,
    title text NOT NULL,
    PRIMARY KEY (id)
);
-- Indexes
CREATE INDEX if not exists groups_idx_title ON dco.groups using hash (title);

-- +goose Down

DROP INDEX IF EXISTS dco.groups_idx_title;
DROP TABLE IF EXISTS dco.groups;
DROP TABLE IF EXISTS dco.clients_segments;
DROP TABLE IF EXISTS dco.segments;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
