-- name: CreateOrder :one
INSERT INTO "Order" (
	"UserUuid") 
VALUES (
    $1
)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM "Order"
WHERE "Uuid" = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM "Order"
ORDER BY "Uuid"
LIMIT $1
OFFSET $2;

-- name: UpdateOrder :one
UPDATE "Order"
  set "UserUuid" = $2
WHERE "Uuid" = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM "Order"
WHERE "Uuid" = $1;