package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/alekseiapa/apple_store/util"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) *User {

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	if err != nil {
		log.Fatal(err)
	}
	arg := CreateUserParams{
		FirstName:      util.RandomUserFirstName(),
		MiddleName:     util.RandomUserMiddleName(),
		LastName:       util.RandomUserLastName(),
		Gender:         util.RandomUserGender(),
		Age:            int16(util.RandomUserAge()),
		Balance:        util.RandomUserBalance(),
		Username:       util.RandomString(12),
		HashedPassword: hashedPassword,
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.MiddleName, user.MiddleName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Gender, user.Gender)
	require.Equal(t, arg.Age, user.Age)
	require.Equal(t, arg.Balance, user.Balance)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.Uuid)
	return &user
}

func createRandomUserWithBalance(t *testing.T, balance float32) *User {
	arg := CreateUserParams{
		FirstName:  util.RandomUserFirstName(),
		MiddleName: util.RandomUserMiddleName(),
		LastName:   util.RandomUserLastName(),
		Gender:     util.RandomUserGender(),
		Age:        int16(util.RandomUserAge()),
		Balance:    balance,
		Username:   util.RandomString(6),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.MiddleName, user.MiddleName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Gender, user.Gender)
	require.Equal(t, arg.Age, user.Age)
	require.Equal(t, arg.Balance, user.Balance)
	require.Equal(t, arg.Username, user.Username)

	require.NotZero(t, user.Uuid)
	return &user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
	createRandomUserWithBalance(t, 1000)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.MiddleName, user2.MiddleName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Gender, user2.Gender)
	require.Equal(t, user1.Age, user2.Age)
	require.Equal(t, user1.Balance, user2.Balance)

}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	balance := util.RandomFloat(10.00, 20.00)
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	if err != nil {
		log.Fatal(err)
	}
	arg := UpdateUserParams{
		Uuid:           user1.Uuid,
		FirstName:      util.RandomString(12),
		MiddleName:     util.RandomString(12),
		LastName:       util.RandomString(12),
		Gender:         util.RandomString(1),
		Age:            int16(util.RandomInt(10, 20)),
		Balance:        balance,
		HashedPassword: hashedPassword,
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user2.FirstName, arg.FirstName)
	require.Equal(t, user2.MiddleName, arg.MiddleName)
	require.Equal(t, user2.LastName, arg.LastName)
	require.Equal(t, user2.Gender, arg.Gender)
	require.Equal(t, user2.Balance, balance)
	require.Equal(t, user2.FullName, fmt.Sprintf("%s %s %s", arg.LastName, arg.FirstName, arg.MiddleName))
	require.Equal(t, user2.Age, arg.Age)
	require.Equal(t, user2.HashedPassword, arg.HashedPassword)

}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	_, err := testQueries.DeleteUser(context.Background(), user1.Uuid)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.Uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)

}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

}
