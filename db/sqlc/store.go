package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provide all functions to execute db queries and transactions
// In order to make a support of transactions we should use the Composition here

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// BuyProductTxParams contains all the necessary parameters to buy a product
type BuyProductTxParams struct {
	UserUuid    int64 `json:"UserUuid"`
	ProductUuid int64 `json:"ProductUuid"`
	Quantity    int32 `json:"Quantity"`
}

// BuyProductTxResult is the result after a successful purchase of a product
type BuyProductTxResult struct {
	User    User    `json:"User"`
	Order   Order   `json:"Order"`
	Product Product `json:"Product"`
}

// Creates on Order.Uuid record, updates User's balance, Add a record to OrderProduct Table
func (store *Store) BuyProductTx(ctx context.Context, arg BuyProductTxParams) (BuyProductTxResult, error) {
	var result BuyProductTxResult

	// we’re accessing the result variable of the outer function from inside this callback function similar for the arg variable.
	//This makes the callback function become a closure. Since Go lacks support for generics type,
	//Closure is often used when we want to get the result from a callback function,
	//because the callback function itself doesn’t know the exact type of the result it should return.
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		product, err := q.GetProductForUpdate(ctx, arg.ProductUuid)
		if err != nil {
			return err
		}
		inStock := product.InStock - arg.Quantity
		if inStock < 0 {
			return fmt.Errorf("sorry you can't buy since there is not enough pcs left. product uuid: %v - %v -> %v pcs left", product.Uuid, product.Description, product.InStock)
		}
		user, err := q.GetUserForUpdate(ctx, arg.UserUuid)
		if err != nil {
			return err
		}
		userBalance := user.Balance - product.Price*float32(arg.Quantity)
		if userBalance < 0 {
			return fmt.Errorf("sorry, you don't have enough money to purchase %v pcs of product uuid: %v - %v", arg.Quantity, product.Uuid, product.Description)
		}
		result.Product, err = q.ReduceProductInStock(ctx, ReduceProductInStockParams{
			Amount: arg.Quantity,
			Uuid:   product.Uuid,
		})
		if err != nil {
			return err
		}
		result.User, err = q.ReduceUserBalance(ctx, ReduceUserBalanceParams{
			Uuid:   user.Uuid,
			Amount: float32(arg.Quantity) * product.Price,
		})
		if err != nil {
			return err
		}
		result.Order, err = q.CreateOrder(ctx, CreateOrderParams{
			UserUuid: user.Uuid,
			Quantity: int64(arg.Quantity),
		})
		if err != nil {
			return err
		}
		_, err = q.CreateOrderProduct(ctx, CreateOrderProductParams{
			OrderUuid:   result.Order.Uuid,
			ProductUuid: result.Product.Uuid,
		})
		if err != nil {
			return err
		}

		return nil

	})

	return result, err
}
