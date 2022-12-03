-- name: CreateUserToUser :one
INSERT INTO "UserToUser" (
	"FirstUserUuid",
    "SecondUserUuid") 
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetUserToUser :one
SELECT * FROM "UserToUser"
WHERE "FirstUserUuid" = $1 
    AND "SecondUserUuid" = $2 
LIMIT 1;

-- name: ListUserToUser :many
SELECT * FROM "UserToUser"
ORDER BY "FirstUserUuid"
LIMIT $1
OFFSET $2;

-- name: UpdateUserToUser :one
UPDATE "UserToUser"
  set "SecondUserUuid" = $3
WHERE "FirstUserUuid" = $1 
    AND "SecondUserUuid" = $2
RETURNING *;