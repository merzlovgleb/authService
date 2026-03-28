-- +goose Up
-- +goose StatementBegin
ALTER TABLE dco.users
add column amount float;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE dco.users
DROP COLUMN IF EXISTS amount;
-- +goose StatementEnd
