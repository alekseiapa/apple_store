package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/alekseiapa/apple_store/db/mock"
	db "github.com/alekseiapa/apple_store/db/sqlc"
	"github.com/alekseiapa/apple_store/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetProductAPI(t *testing.T) {
	product := randomProduct()

	testCases := []struct {
		name          string
		ProductUuid   int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			ProductUuid: product.Uuid,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.Uuid)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProduct(t, recorder.Body, product)
			},
		},
		{
			name:        "NotFound",
			ProductUuid: product.Uuid,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.Uuid)).
					Times(1).
					Return(db.Product{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:        "InternalError",
			ProductUuid: product.Uuid,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.Uuid)).
					Times(1).
					Return(db.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:        "InvalidID",
			ProductUuid: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewTestServer(t, store)
			//For testing an HTTP API in Go, we don’t have to start a real HTTP server,
			//Instead, we can just use the Recorder feature of the httptest package
			//to record the response of the API request.
			//So here we call httptest.NewRecorder() to create a new ResponseRecorder
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/api/products/%d?currency=USD", tc.ProductUuid)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func randomProduct() db.Product {
	return db.Product{
		Uuid:        int64(util.RandomInt(1, 1000)),
		Description: util.RandomProductDescription(),
		Price:       util.RandomProductPrice(),
		InStock:     util.RandomProductInStock(),
	}
}

func requireBodyMatchProduct(t *testing.T, body *bytes.Buffer, product db.Product) {
	data, err := ioutil.ReadAll(body)

	if err != nil {
		log.Fatal(err)
	}

	var gotProduct productResponse
	err = json.Unmarshal(data, &gotProduct)
	fmt.Println(gotProduct)
	require.NoError(t, err)
	require.Equal(t, product.Uuid, gotProduct.Uuid)
	require.Equal(t, product.Price, gotProduct.Price)
	require.Equal(t, product.Description, gotProduct.Description)
	require.Equal(t, product.InStock, gotProduct.InStock)
}
