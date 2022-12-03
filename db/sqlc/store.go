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
	User    User    `json:"User"`
	Product Product `json:"Product"`
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

		inStock := arg.Product.InStock - 1
		if inStock < 0 {
			return fmt.Errorf("product uuid: %v - %v is out of stock", arg.Product.Uuid, arg.Product.Description)
		}
		userBalance := arg.User.Balance - int64(arg.Product.Price)
		if userBalance < 0 {
			return fmt.Errorf("sorry, you don't have enough balance to purchase product uuid: %v - %v", arg.Product.Uuid, arg.Product.Description)
		}
		result.Product, err = q.UpdateProduct(ctx, UpdateProductParams{
			Uuid:        arg.Product.Uuid,
			Description: arg.Product.Description,
			Price:       arg.Product.Price,
			InStock:     inStock,
		})
		if err != nil {
			return err
		}
		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			Uuid:       arg.User.Uuid,
			FirstName:  arg.User.FirstName,
			MiddleName: arg.User.MiddleName,
			LastName:   arg.User.LastName,
			Gender:     arg.User.Gender,
			Age:        arg.User.Age,
			Balance:    userBalance,
		})
		if err != nil {
			return err
		}
		result.Order, err = q.CreateOrder(ctx, arg.User.Uuid)
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
