-- name: CreateUser :one
INSERT INTO "User" (
	"FirstName", 
	"MiddleName", 
	"LastName", 
	"Gender", 
	"Age",
  "Balance",
  "Username",
  "HashedPassword") 
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "User"
WHERE "Uuid" = $1 LIMIT 1;

-- name: GetUserByUserName :one
SELECT * FROM "User"
WHERE "Username" = $1 LIMIT 1;

-- This will allow us to block transactions till the end of commit
-- name: GetUserForUpdate :one
SELECT * FROM "User"
WHERE "Uuid" = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: ReduceUserBalance :one
UPDATE "User"
  set "Balance" = "Balance" - sqlc.arg(amount) 
WHERE "Uuid" = sqlc.arg(Uuid)
RETURNING *;


-- name: ListUsers :many
SELECT * FROM "User"
ORDER BY "Uuid" ASC
LIMIT $1
OFFSET $2;


-- name: UpdateUser :one
UPDATE "User"
  set "FirstName" = $2,
      "MiddleName" = $3,
      "LastName" = $4,
      "Gender" = $5,
      "Age" = $6,
      "Balance" = $7,
      "HashedPassword" = $8
WHERE "Uuid" = $1
RETURNING *;

-- name: DeleteUser :execrows
DELETE FROM "User"
WHERE "Uuid" = $1;