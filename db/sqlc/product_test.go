package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/alekseiapa/apple_store/util"

	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Description: util.RandomProductDescription(),
		Price:       util.RandomProductPrice(),
		InStock:     util.RandomProductInStock(),
	}
	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Description, product.Description)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.InStock, product.InStock)

	require.NotZero(t, product.Uuid)
	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	product2, err := testQueries.GetProduct(context.Background(), product1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.Description, product2.Description)
	require.Equal(t, product1.Price, product2.Price)
	require.Equal(t, product1.InStock, product2.InStock)

}

func TestUpdateProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	price := util.RandomProductPrice()
	inStock := util.RandomProductInStock()
	arg := UpdateProductParams{
		Uuid:        product1.Uuid,
		Description: util.RandomString(12),
		Price:       price,
		InStock:     inStock,
	}
	product2, err := testQueries.UpdateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product2.Description, arg.Description)
	require.Equal(t, product2.Price, arg.Price)
	require.Equal(t, product2.InStock, arg.InStock)

}

func TestDeleteProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	err := testQueries.DeleteProduct(context.Background(), product1.Uuid)
	require.NoError(t, err)

	product2, err := testQueries.GetUser(context.Background(), product1.Uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)

}

func TestListProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	arg := ListProductsParams{
		Limit:  5,
		Offset: 5,
	}

	products, err := testQueries.ListProducts(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, products, 5)

	for _, product := range products {
		require.NotEmpty(t, product)
	}

}
