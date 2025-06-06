// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: go_db_user.sql

package database

import (
	"context"
	"database/sql"
)

const checkEmailExists = `-- name: CheckEmailExists :one
SELECT EXISTS(SELECT 1 FROM go_db_user WHERE email = ?) AS email_exists
`

func (q *Queries) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkEmailExists, email)
	var email_exists bool
	err := row.Scan(&email_exists)
	return email_exists, err
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO go_db_user (firstname, lastname, username, email, password_hash, password_salt, status, create_date, edit_date) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
`

type CreateUserParams struct {
	Firstname    string
	Lastname     string
	Username     string
	Email        string
	PasswordHash string
	PasswordSalt string
	Status       string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.Firstname,
		arg.Lastname,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
		arg.PasswordSalt,
		arg.Status,
	)
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE go_db_user SET disabled = 1 WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id uint32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, firstname, lastname, username, email, password_hash, disabled, status, last_login_date, create_date, edit_date, password_salt, is_2fa_enabled FROM go_db_user WHERE email = ? LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GoDbUser, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i GoDbUser
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.Disabled,
		&i.Status,
		&i.LastLoginDate,
		&i.CreateDate,
		&i.EditDate,
		&i.PasswordSalt,
		&i.Is2faEnabled,
	)
	return i, err
}

const getUserInfo = `-- name: GetUserInfo :one
SELECT id, firstname, lastname, username, email, password_hash, disabled, status, last_login_date, create_date, edit_date, password_salt, is_2fa_enabled FROM go_db_user WHERE id = ? LIMIT 1
`

func (q *Queries) GetUserInfo(ctx context.Context, id uint32) (GoDbUser, error) {
	row := q.db.QueryRowContext(ctx, getUserInfo, id)
	var i GoDbUser
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.Disabled,
		&i.Status,
		&i.LastLoginDate,
		&i.CreateDate,
		&i.EditDate,
		&i.PasswordSalt,
		&i.Is2faEnabled,
	)
	return i, err
}

const updateUserLastLoginDate = `-- name: UpdateUserLastLoginDate :exec
UPDATE go_db_user SET last_login_date = NOW() WHERE id = ?
`

func (q *Queries) UpdateUserLastLoginDate(ctx context.Context, id uint32) error {
	_, err := q.db.ExecContext(ctx, updateUserLastLoginDate, id)
	return err
}

const updateUserPasswordHash = `-- name: UpdateUserPasswordHash :exec
UPDATE go_db_user SET password_hash = ? WHERE id = ?
`

type UpdateUserPasswordHashParams struct {
	PasswordHash string
	ID           uint32
}

func (q *Queries) UpdateUserPasswordHash(ctx context.Context, arg UpdateUserPasswordHashParams) error {
	_, err := q.db.ExecContext(ctx, updateUserPasswordHash, arg.PasswordHash, arg.ID)
	return err
}

const updateUserStatusByEmail = `-- name: UpdateUserStatusByEmail :exec
UPDATE go_db_user SET status = ?, edit_date = NOW() WHERE email = ?
`

type UpdateUserStatusByEmailParams struct {
	Status string
	Email  string
}

func (q *Queries) UpdateUserStatusByEmail(ctx context.Context, arg UpdateUserStatusByEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateUserStatusByEmail, arg.Status, arg.Email)
	return err
}
