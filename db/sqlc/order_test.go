package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/alekseiapa/apple_store/util"

	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T) Order {
	arg := CreateOrderParams{
		UserUuid: util.RandomOrderUseruuid(),
		Quantity: util.RandomOrderQuantity(),
	}
	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.UserUuid, order.UserUuid)
	require.Equal(t, arg.Quantity, order.Quantity)
	require.NotZero(t, order.Uuid)
	return order
}
func TestCreateOrder(t *testing.T) {
	createRandomOrder(t)
}

func TestGetOrder(t *testing.T) {
	order1 := createRandomOrder(t)
	order2, err := testQueries.GetOrder(context.Background(), order1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.UserUuid, order2.UserUuid)
	require.Equal(t, order1.Quantity, order2.Quantity)
}

func TestUpdateOrder(t *testing.T) {
	order1 := createRandomOrder(t)
	arg := UpdateOrderParams{
		Uuid:     order1.Uuid,
		UserUuid: int64(util.RandomInt(10, 20)),
		Quantity: util.RandomOrderQuantity(),
	}
	order2, err := testQueries.UpdateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order2.Uuid, arg.Uuid)
	require.Equal(t, order2.UserUuid, arg.UserUuid)
	require.Equal(t, order2.Quantity, arg.Quantity)

}

func TestDeleteOrder(t *testing.T) {
	order1 := createRandomOrder(t)
	_, err := testQueries.DeleteOrder(context.Background(), order1.Uuid)
	require.NoError(t, err)

	order2, err := testQueries.GetOrder(context.Background(), order1.Uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, order2)

}

func TestListOrders(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrder(t)
	}

	arg := ListOrdersParams{
		Limit:  5,
		Offset: 5,
	}

	orders, err := testQueries.ListOrders(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, orders, 5)

	for _, order := range orders {
		require.NotEmpty(t, order)
	}

}
