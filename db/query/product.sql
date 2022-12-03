-- name: CreateProduct :one
INSERT INTO "Product" (
    "Description",
    "Price",
    "InStock") 
VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM "Product"
WHERE "Uuid" = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM "Product"
ORDER BY "Uuid"
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :one
UPDATE "Product"
    set "Description" = $2,
        "Price" = $3,
        "InStock" = $4
WHERE "Uuid" = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM "Product"
WHERE "Uuid" = $1;