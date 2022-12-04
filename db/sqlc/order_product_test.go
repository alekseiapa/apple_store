package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomOrderProduct(t *testing.T) OrderProduct {
	order := createRandomOrder(t)
	product := createRandomProduct(t)
	arg := CreateOrderProductParams{
		OrderUuid:   order.Uuid,
		ProductUuid: product.Uuid,
	}
	orderProduct, err := testQueries.CreateOrderProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderProduct)

	require.Equal(t, arg.OrderUuid, orderProduct.OrderUuid)
	require.Equal(t, arg.ProductUuid, orderProduct.ProductUuid)

	return orderProduct
}
func TestCreateOrderProduct(t *testing.T) {
	createRandomOrderProduct(t)
}

func TestGetOrderProduct(t *testing.T) {
	orderProduct1 := createRandomOrderProduct(t)
	arg := GetOrderProductParams(orderProduct1)
	orderProduct2, err := testQueries.GetOrderProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderProduct2)

	require.Equal(t, orderProduct1.OrderUuid, orderProduct2.OrderUuid)
	require.Equal(t, orderProduct1.ProductUuid, orderProduct2.ProductUuid)

}

func TestUpdateOrderProduct(t *testing.T) {
	orderProduct1 := createRandomOrderProduct(t)

	newProduct := createRandomProduct(t)

	arg := UpdateOrderProductParams{
		OrderUuid:     orderProduct1.OrderUuid,
		ProductUuid:   orderProduct1.ProductUuid,
		ProductUuid_2: newProduct.Uuid,
	}
	orderProduct2, err := testQueries.UpdateOrderProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderProduct2)

	require.Equal(t, orderProduct2.OrderUuid, arg.OrderUuid)
	require.Equal(t, orderProduct2.ProductUuid, arg.ProductUuid_2)

}

func TestDeleteOrderProduct(t *testing.T) {
	orderProduct1 := createRandomOrderProduct(t)

	_, err := testQueries.DeleteOrder(context.Background(), orderProduct1.OrderUuid)

	require.NoError(t, err)

	getOrderProductArg := GetOrderProductParams(orderProduct1)
	orderProduct2, err := testQueries.GetOrderProduct(context.Background(), getOrderProductArg)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, orderProduct2)

}

func TestOrderProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrderProduct(t)
	}

	arg := ListOrderProductsParams{
		Limit:  5,
		Offset: 5,
	}

	orderProducts, err := testQueries.ListOrderProducts(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, orderProducts, 5)

	for _, orderProduct := range orderProducts {
		require.NotEmpty(t, orderProduct)
	}

}
