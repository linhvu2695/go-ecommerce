-- +goose Up
-- +goose StatementBegin
ALTER TABLE `go_db_user`
ADD COLUMN `is_2fa_enabled` BOOLEAN DEFAULT FALSE ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `go_db_user`
DROP COLUMN `is_2fa_enabled`;
-- +goose StatementEnd
