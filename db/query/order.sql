-- name: CreateOrder :one
INSERT INTO "Order" (
	"UserUuid",
  "Quantity") 
VALUES (
    $1, $2
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
  set "UserUuid" = $2,
      "Quantity" = $3
WHERE "Uuid" = $1
RETURNING *;

-- name: DeleteOrder :execrows
DELETE FROM "Order"
WHERE "Uuid" = $1;

-- TODO: Read the doc of sqlc to determine the way how to retrieve many2many data
-- SELECT 
-- "Order"."Uuid" AS "Order Uuid",
-- (SELECT "User"."FullName" FROM "User" WHERE "Order"."UserUuid" = "User"."Uuid") AS "User Name",
-- "Order"."Quantity" AS "Order Qty", 
-- "Product"."Uuid" AS "Product Uuid",
-- "Product"."Description" AS "Product Description"
-- FROM "OrderProduct"
-- INNER JOIN "Order" ON "Order"."Uuid" = "OrderProduct"."OrderUuid"
-- INNER JOIN "Product" ON "Product"."Uuid" = "OrderProduct"."ProductUuid"
-- WHERE "Order"."Uuid" = 1;