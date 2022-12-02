// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "User" (
	"FirstName", 
	"MiddleName", 
	"LastName", 
	"Gender", 
	"Age") 
VALUES (
    $1, $2, $3, $4, $5
)
RETURNING "Uuid", "FirstName", "MiddleName", "LastName", "FullName", "Gender", "Age"
`

type CreateUserParams struct {
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName   string `json:"LastName"`
	Gender     string `json:"Gender"`
	Age        int16  `json:"Age"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.MiddleName,
		arg.LastName,
		arg.Gender,
		arg.Age,
	)
	var i User
	err := row.Scan(
		&i.Uuid,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.FullName,
		&i.Gender,
		&i.Age,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "User"
WHERE "Uuid" = $1
`

func (q *Queries) DeleteUser(ctx context.Context, uuid int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, uuid)
	return err
}

const getUser = `-- name: GetUser :one
SELECT "Uuid", "FirstName", "MiddleName", "LastName", "FullName", "Gender", "Age" FROM "User"
WHERE "Uuid" = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, uuid int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, uuid)
	var i User
	err := row.Scan(
		&i.Uuid,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.FullName,
		&i.Gender,
		&i.Age,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT "Uuid", "FirstName", "MiddleName", "LastName", "FullName", "Gender", "Age" FROM "User"
ORDER BY "FullName"
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Uuid,
			&i.FirstName,
			&i.MiddleName,
			&i.LastName,
			&i.FullName,
			&i.Gender,
			&i.Age,
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

const updateUser = `-- name: UpdateUser :one
UPDATE "User"
  set "FirstName" = $2,
      "MiddleName" = $3,
      "LastName" = $4,
      "Gender" = $5,
      "Age" = $6
WHERE "Uuid" = $1
RETURNING "Uuid", "FirstName", "MiddleName", "LastName", "FullName", "Gender", "Age"
`

type UpdateUserParams struct {
	Uuid       int64  `json:"Uuid"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName   string `json:"LastName"`
	Gender     string `json:"Gender"`
	Age        int16  `json:"Age"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Uuid,
		arg.FirstName,
		arg.MiddleName,
		arg.LastName,
		arg.Gender,
		arg.Age,
	)
	var i User
	err := row.Scan(
		&i.Uuid,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.FullName,
		&i.Gender,
		&i.Age,
	)
	return i, err
}
