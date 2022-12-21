// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "Users" (
    email,
    username,
    password
    ) VALUES (
    $1, $2, $3
) RETURNING id, email, username, password
`

type CreateUserParams struct {
	Email    sql.NullString `json:"email"`
	Username sql.NullString `json:"username"`
	Password sql.NullString `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Username, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.Password,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "Users" WHERE "id" = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, email, username, password FROM "Users"
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Username,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOneUser = `-- name: GetOneUser :one
SELECT id, email, username, password FROM "Users" WHERE "Users"."id" = $1 LIMIT 1
`

func (q *Queries) GetOneUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getOneUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.Password,
	)
	return i, err
}

const updateEmail = `-- name: UpdateEmail :exec
UPDATE "Users" SET "email" = $2 WHERE "id" = $1
`

type UpdateEmailParams struct {
	ID    int32          `json:"id"`
	Email sql.NullString `json:"email"`
}

func (q *Queries) UpdateEmail(ctx context.Context, arg UpdateEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateEmail, arg.ID, arg.Email)
	return err
}
