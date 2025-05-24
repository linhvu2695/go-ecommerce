-- name: GetUserByEmail :one
SELECT * FROM go_db_user WHERE email = ? LIMIT 1;

-- name: GetUserInfo :one
SELECT * FROM go_db_user WHERE id = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO go_db_user (firstname, lastname, username, email, password_hash, password_salt, status, create_date, edit_date) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW());

-- name: UpdateUserStatusByEmail :exec
UPDATE go_db_user SET status = ?, edit_date = NOW() WHERE email = ?;

-- name: UpdateUserLastLoginDate :exec
UPDATE go_db_user SET last_login_date = NOW() WHERE id = ?;

-- name: UpdateUserPasswordHash :exec
UPDATE go_db_user SET password_hash = ? WHERE id = ?;

-- name: DeleteUser :exec
UPDATE go_db_user SET disabled = 1 WHERE id = ?;

-- name: CheckEmailExists :one
SELECT EXISTS(SELECT 1 FROM go_db_user WHERE email = ?) AS email_exists;