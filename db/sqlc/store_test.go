package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuyTx(t *testing.T) {

	store := NewStore(testDB)

	user := createRandomUser(t)
	user.Balance = 1000

	fmt.Println(">> User balance before:", user.Balance)

	// So the best way to make sure that our purchase works well
	// is to run it with several concurrent go routines.
	// Let’s say I want to run n = 5 concurrent buy purchases
	// And each of them will reduce a price of product of 10 from user_1.
	// So I will use a simple for loop with n iterations
	// And inside the loop,
	// we use the go keyword to start a new routine.
	// run n concurrent purchases

	errs := make(chan error)
	results := make(chan BuyProductTxResult)

	n := 5

	// run n concurrent purchases
	for i := 0; i < n; i++ {
		product := createRandomProduct(t)
		product.Price = 100
		go func() {
			result, err := store.BuyProductTx(context.Background(), BuyProductTxParams{
				User:    user,
				Product: product,
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

	}

}
