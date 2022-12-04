package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuyTx(t *testing.T) {

	store := NewStore(testDB)

	user := createRandomUserWithBalance(t, 1000)
	product := createRandomProductWithPriceAndInStock(t, 100, 6)

	// So the best way to make sure that our purchase works well
	// is to run it with several concurrent go routines.
	// Let’s say I want to run n = 5 concurrent buy purchases
	// And each of them will reduce a price of product of 100 from user.
	// So I will use a simple for loop with n iterations
	// And inside the loop,
	// we use the go keyword to start a new routine.
	// run n concurrent purchases

	errs := make(chan error)
	results := make(chan BuyProductTxResult)

	var (
		finalBalance  float32
		finalInStock  int32
		toBuyPcs      int32
		totalToBuyPcs int32
	)

	n := 5
	toBuyPcs = 1
	totalToBuyPcs = toBuyPcs * int32(n)

	// run n concurrent purchases
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.BuyProductTx(context.Background(), BuyProductTxParams{
				UserUuid:    user.Uuid,
				ProductUuid: product.Uuid,
				Quantity:    toBuyPcs,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check User
		userDB := result.User
		require.NotEmpty(t, userDB)
		require.Equal(t, user.Uuid, userDB.Uuid)
		require.Equal(t, user.FirstName, userDB.FirstName)
		require.Equal(t, user.MiddleName, userDB.MiddleName)
		require.Equal(t, user.LastName, userDB.LastName)
		require.Equal(t, user.Gender, userDB.Gender)
		require.Equal(t, user.Age, userDB.Age)

		require.NotZero(t, userDB.Balance)

		// check Order
		orderDB := result.Order
		require.NotEmpty(t, orderDB)

		// check Product
		productDB := result.Product
		require.NotEmpty(t, productDB)
		finalInStock = result.Product.InStock
		finalBalance = result.User.Balance

	}
	require.Equal(t, product.InStock-totalToBuyPcs, finalInStock)
	require.Equal(t, user.Balance-float32(totalToBuyPcs)*product.Price, finalBalance)

}

func TestBuyNotEnoughInStockTx(t *testing.T) {

	store := NewStore(testDB)

	user := createRandomUserWithBalance(t, 1000)
	product := createRandomProductWithPriceAndInStock(t, 100, 1)

	errs := make(chan error)

	var (
		toBuyPcs int32
	)

	n := 5
	toBuyPcs = 5

	// run n concurrent purchases
	for i := 0; i < n; i++ {
		go func() {
			_, err := store.BuyProductTx(context.Background(), BuyProductTxParams{
				UserUuid:    user.Uuid,
				ProductUuid: product.Uuid,
				Quantity:    toBuyPcs,
			})
			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.ErrorContains(t, err, "sorry you can't buy since there is not enough pcs left")
	}

}

func TestBuyNotEnoughMoneyTx(t *testing.T) {

	store := NewStore(testDB)

	user := createRandomUserWithBalance(t, 100)
	product := createRandomProductWithPriceAndInStock(t, 100, 10)

	errs := make(chan error)

	var (
		toBuyPcs int32
	)

	n := 5
	toBuyPcs = 5

	// run n concurrent purchases
	for i := 0; i < n; i++ {
		go func() {
			_, err := store.BuyProductTx(context.Background(), BuyProductTxParams{
				UserUuid:    user.Uuid,
				ProductUuid: product.Uuid,
				Quantity:    toBuyPcs,
			})
			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.ErrorContains(t, err, "sorry, you don't have enough money to purchase")
	}

}
