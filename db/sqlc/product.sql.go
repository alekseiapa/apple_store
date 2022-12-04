// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: product.sql

package db

import (
	"context"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO "Product" (
    "Description",
    "Price",
    "InStock") 
VALUES (
    $1, $2, $3
)
RETURNING "Uuid", "Description", "Price", "InStock"
`

type CreateProductParams struct {
	Description string  `json:"Description"`
	Price       float32 `json:"Price"`
	InStock     int32   `json:"InStock"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct, arg.Description, arg.Price, arg.InStock)
	var i Product
	err := row.Scan(
		&i.Uuid,
		&i.Description,
		&i.Price,
		&i.InStock,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :execrows
DELETE FROM "Product"
WHERE "Uuid" = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, uuid int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteProduct, uuid)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getProduct = `-- name: GetProduct :one
SELECT "Uuid", "Description", "Price", "InStock" FROM "Product"
WHERE "Uuid" = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, uuid int64) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, uuid)
	var i Product
	err := row.Scan(
		&i.Uuid,
		&i.Description,
		&i.Price,
		&i.InStock,
	)
	return i, err
}

const getProductForUpdate = `-- name: GetProductForUpdate :one
SELECT "Uuid", "Description", "Price", "InStock" FROM "Product"
WHERE "Uuid" = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetProductForUpdate(ctx context.Context, uuid int64) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductForUpdate, uuid)
	var i Product
	err := row.Scan(
		&i.Uuid,
		&i.Description,
		&i.Price,
		&i.InStock,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT "Uuid", "Description", "Price", "InStock" FROM "Product"
ORDER BY "Uuid"
LIMIT $1
OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.Uuid,
			&i.Description,
			&i.Price,
			&i.InStock,
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

const reduceProductInStock = `-- name: ReduceProductInStock :one
UPDATE "Product"
  set "InStock" = "InStock" - $1 
WHERE "Uuid" = $2
RETURNING "Uuid", "Description", "Price", "InStock"
`

type ReduceProductInStockParams struct {
	Amount int32 `json:"amount"`
	Uuid   int64 `json:"uuid"`
}

func (q *Queries) ReduceProductInStock(ctx context.Context, arg ReduceProductInStockParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, reduceProductInStock, arg.Amount, arg.Uuid)
	var i Product
	err := row.Scan(
		&i.Uuid,
		&i.Description,
		&i.Price,
		&i.InStock,
	)
	return i, err
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE "Product"
    set "Description" = $2,
        "Price" = $3,
        "InStock" = $4
WHERE "Uuid" = $1
RETURNING "Uuid", "Description", "Price", "InStock"
`

type UpdateProductParams struct {
	Uuid        int64   `json:"Uuid"`
	Description string  `json:"Description"`
	Price       float32 `json:"Price"`
	InStock     int32   `json:"InStock"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.Uuid,
		arg.Description,
		arg.Price,
		arg.InStock,
	)
	var i Product
	err := row.Scan(
		&i.Uuid,
		&i.Description,
		&i.Price,
		&i.InStock,
	)
	return i, err
}
