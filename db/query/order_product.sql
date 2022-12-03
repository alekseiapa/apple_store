-- name: CreateOrderProduct :one
INSERT INTO "OrderProduct" (
	"OrderUuid",
    "ProductUuid") 
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetOrderProduct :one
SELECT * FROM "OrderProduct"
WHERE "OrderUuid" = $1 
    AND "ProductUuid" = $2 
LIMIT 1;

-- name: ListOrderProducts :many
SELECT * FROM "OrderProduct"
ORDER BY "OrderUuid"
LIMIT $1
OFFSET $2;

-- name: UpdateOrderProduct :one
UPDATE "OrderProduct"
  set "ProductUuid" = $3
WHERE "OrderUuid" = $1 
    AND "ProductUuid" = $2
RETURNING *;