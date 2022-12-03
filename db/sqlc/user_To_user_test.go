package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUserToUser(t *testing.T) UserToUser {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	arg := CreateUserToUserParams{
		FirstUserUuid:  user1.Uuid,
		SecondUserUuid: user2.Uuid,
	}
	userToUser, err := testQueries.CreateUserToUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userToUser)

	require.Equal(t, arg.FirstUserUuid, userToUser.FirstUserUuid)
	require.Equal(t, arg.SecondUserUuid, userToUser.SecondUserUuid)

	return userToUser
}

func TestCreateUserToUser(t *testing.T) {
	createRandomUserToUser(t)
}

func TestGetUserToUser(t *testing.T) {
	userToUser1 := createRandomUserToUser(t)
	arg := GetUserToUserParams(userToUser1)
	userToUser2, err := testQueries.GetUserToUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userToUser2)

	require.Equal(t, userToUser1.FirstUserUuid, userToUser2.FirstUserUuid)
	require.Equal(t, userToUser1.SecondUserUuid, userToUser2.SecondUserUuid)

}

func TestUpdateUserToUser(t *testing.T) {
	userToUser1 := createRandomUserToUser(t)
	newUser := createRandomUser(t)

	arg := UpdateUserToUserParams{
		FirstUserUuid:    userToUser1.FirstUserUuid,
		SecondUserUuid:   userToUser1.SecondUserUuid,
		SecondUserUuid_2: newUser.Uuid,
	}
	userToUser2, err := testQueries.UpdateUserToUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userToUser2)

	require.Equal(t, userToUser2.FirstUserUuid, arg.FirstUserUuid)
	require.Equal(t, userToUser2.SecondUserUuid, arg.SecondUserUuid_2)

}

func TestDeleteUserToUser(t *testing.T) {
	userToUser1 := createRandomUserToUser(t)

	err := testQueries.DeleteUser(context.Background(), userToUser1.FirstUserUuid)

	require.NoError(t, err)

	getUserToUserArg := GetUserToUserParams(userToUser1)
	userToUser2, err := testQueries.GetUserToUser(context.Background(), getUserToUserArg)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, userToUser2)

}

func TestListUserToUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUserToUser(t)
	}

	arg := ListUserToUserParams{
		Limit:  5,
		Offset: 5,
	}

	userToUsers, err := testQueries.ListUserToUser(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, userToUsers, 5)

	for _, userToUser := range userToUsers {
		require.NotEmpty(t, userToUser)
	}

}
