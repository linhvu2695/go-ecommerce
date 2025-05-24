-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `go_db_user` (
    id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    disabled BOOLEAN DEFAULT FALSE,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    last_login_date TIMESTAMP NULL DEFAULT NULL,
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    edit_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    KEY `idx_username` (username),
    KEY `idx_email` (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `go_db_user`;
-- +goose StatementEnd
