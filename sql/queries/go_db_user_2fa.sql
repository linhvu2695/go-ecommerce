-- name: Enable2FAEmail :exec
INSERT INTO go_db_user_2fa (user_id, type, email, secret, is_active, create_date, edit_date) 
VALUES (?, ?, ?, "SECRET", FALSE, NOW(), NOW());

-- name: Disable2FA :exec
UPDATE go_db_user_2fa SET is_active = FALSE, edit_date = NOW() 
WHERE user_id = ? AND type = ?;

-- name: Update2FAStatus :exec
UPDATE go_db_user_2fa SET is_active = TRUE, edit_date = NOW() 
WHERE user_id = ? AND type = ? AND is_active = FALSE;

-- name: Verify2FA :one
SELECT COUNT(*) FROM go_db_user_2fa
WHERE user_id = ? AND type = ? AND is_active = TRUE;

-- name: Get2FAStatus :one
SELECT is_active FROM go_db_user_2fa
WHERE user_id = ? AND type = ?;

-- name: Is2FAEnabled :one
SELECT COUNT(*) FROM go_db_user_2fa
WHERE user_id = ? AND is_active = TRUE;

-- name: AddOrUpdatePhone :exec
INSERT INTO go_db_user_2fa (user_id, phone, is_active) 
VALUES (?, ?, TRUE)
ON DUPLICATE KEY UPDATE phone = ?, edit_date = NOW();

-- name: AddOrUpdateEmail :exec
INSERT INTO go_db_user_2fa (user_id, email, is_active) 
VALUES (?, ?, TRUE)
ON DUPLICATE KEY UPDATE email = ?, edit_date = NOW();

-- name: GetUser2FAMethods :many
SELECT * FROM go_db_user_2fa
WHERE user_id = ?;

-- name: Reactivate2FA :exec
UPDATE go_db_user_2fa SET is_active = TRUE, edit_date = NOW() 
WHERE user_id = ? AND type = ?;

-- name: Remove2FA :exec
DELETE FROM go_db_user_2fa
WHERE user_id = ? AND type = ?;

-- name: CountActive2FAMethods :one
SELECT COUNT(*) FROM go_db_user_2fa
WHERE user_id = ? AND is_active = TRUE;

-- name: Get2FAById :one
SELECT * FROM go_db_user_2fa WHERE id = ?;

-- name: Get2FAByUserAndType :one
SELECT * FROM go_db_user_2fa WHERE user_id = ? AND type = ?;
