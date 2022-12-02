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


-- name: UpdateUser :one
UPDATE "User"
  set "FirstName" = $2,
      "MiddleName" = $3,
      "LastName" = $4,
      "Gender" = $5,
      "Age" = $6
WHERE "Uuid" = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "User"
WHERE "Uuid" = $1;