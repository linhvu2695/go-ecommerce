-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `go_db_user_2fa` (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id INT UNSIGNED NOT NULL,
    type ENUM('SMS', 'EMAIL', 'APP') NOT NULL,
    secret VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NULL,
    email VARCHAR(255) NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    edit_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id),

    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user_2fa';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `go_db_user_2fa`;
-- +goose StatementEnd
