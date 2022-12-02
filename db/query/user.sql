-- name: CreateUser :one
INSERT INTO "User" (
	"FirstName", 
	"MiddleName", 
	"LastName", 
	"Gender", 
	"Age") 
VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "User"
WHERE "Uuid" = $1 LIMIT 1;


-- name: ListUsers :many
SELECT * FROM "User"
ORDER BY "FullName"
LIMIT $1
OFFSET $2;


-- name: UpdateUserFirstName :one
UPDATE "User"
  set "FirstName" = $2
WHERE "Uuid" = $1
RETURNING *;

-- name: UpdateUserMiddleName :one
UPDATE "User"
  set "MiddleName" = $2
WHERE "Uuid" = $1
RETURNING *;

-- name: UpdateUserLastName :one
UPDATE "User"
  set "LastName" = $2
WHERE "Uuid" = $1
RETURNING *;

-- name: UpdateUserGender :one
UPDATE "User"
  set "Gender" = $2
WHERE "Uuid" = $1
RETURNING *;

-- name: UpdateUserAge :one
UPDATE "User"
  set "Age" = $2
WHERE "Uuid" = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "User"
WHERE "Uuid" = $1;