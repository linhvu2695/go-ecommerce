-- +goose Up
-- +goose StatementBegin
ALTER TABLE `go_db_user`
ADD COLUMN `password_salt` VARCHAR(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `go_db_user`
DROP COLUMN `password_salt`;
-- +goose StatementEnd
