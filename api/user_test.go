package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/alekseiapa/apple_store/db/mock"
	db "github.com/alekseiapa/apple_store/db/sqlc"
	"github.com/alekseiapa/apple_store/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestGetUserAPI(t *testing.T) {
	user, _ := randomUser(t)

	testCases := []struct {
		name          string
		UserUuid      int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:     "OK",
			UserUuid: user.Uuid,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Uuid)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name:     "NotFound",
			UserUuid: user.Uuid,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Uuid)).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:     "InternalError",
			UserUuid: user.Uuid,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Uuid)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:     "InvalidID",
			UserUuid: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
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
			url := fmt.Sprintf("/api/users/%d", tc.UserUuid)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUser(t)
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"first_name":  user.FirstName,
				"middle_name": user.MiddleName,
				"last_name":   user.LastName,
				"gender":      user.Gender,
				"age":         user.Age,
				"balance":     user.Balance,
				"username":    user.Username,
				"password":    password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"first_name":  user.FirstName,
				"middle_name": user.MiddleName,
				"last_name":   user.LastName,
				"gender":      user.Gender,
				"age":         user.Age,
				"balance":     user.Balance,
				"username":    user.Username,
				"password":    password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "DuplicateUsername",
			body: gin.H{
				"first_name":  user.FirstName,
				"middle_name": user.MiddleName,
				"last_name":   user.LastName,
				"gender":      user.Gender,
				"age":         user.Age,
				"balance":     user.Balance,
				"username":    user.Username,
				"password":    password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "InvalidUsername",
			body: gin.H{
				"first_name":  user.FirstName,
				"middle_name": user.MiddleName,
				"last_name":   user.LastName,
				"gender":      user.Gender,
				"age":         user.Age,
				"balance":     user.Balance,
				"username":    "invalid-username",
				"password":    password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TooShortPassword",
			body: gin.H{
				"first_name":  user.FirstName,
				"middle_name": user.MiddleName,
				"last_name":   user.LastName,
				"gender":      user.Gender,
				"age":         user.Age,
				"balance":     user.Balance,
				"username":    user.Username,
				"password":    "123",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	// We iterate through all of these cases,
	// and run a separate sub-test for each of them.
	// In each sub-test, we create a new gomock controller,
	// and use it to build a new mock DB store.
	// Then we call the buildStubs() function of the test case to setup the stubs for that store.
	// After that, we create a new server using the mock store,
	// and create a new HTTP response recorder to record the result of the API call.
	// Next we marshal the input request body to JSON,
	// And make a new POST request to the get-user API endpoint with that JSON data.
	// We call server.router.ServeHTTP function with the recorder and request object,
	// And finally just call tc.checkResponse function to check the result.

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/api/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	user = db.User{
		Uuid:           int64(util.RandomInt(1, 1000)),
		FirstName:      util.RandomUserFirstName(),
		MiddleName:     util.RandomUserMiddleName(),
		LastName:       util.RandomUserLastName(),
		Gender:         "M",
		Age:            int16(util.RandomUserAge()),
		Balance:        util.RandomUserBalance(),
		HashedPassword: hashedPassword,
		Username:       util.RandomString(6),
	}

	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser userResponse
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)

	require.Equal(t, user.LastName, gotUser.LastName)
	require.Equal(t, user.MiddleName, gotUser.MiddleName)
	require.Equal(t, user.Gender, gotUser.Gender)
	require.Equal(t, user.Age, gotUser.Age)
	require.Equal(t, user.Balance, gotUser.Balance)
	require.Equal(t, user.Username, gotUser.Username)
}
