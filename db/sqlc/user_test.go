package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/levietcuong2602/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	password := utils.RandomString(16)
	hashedPassword, err := utils.EncryptPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, password)

	arg := CreateUserParams{
		Username:       fmt.Sprintf("%s@gmail.com", utils.RandomString(6)),
		HashedPassword: hashedPassword,
		Email:          fmt.Sprintf("%s@gmail.com", utils.RandomString(6)),
		FullName:       utils.RandomString(8),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)

	err = utils.ComparePassword(password, user.HashedPassword)
	require.NoError(t, err)

	require.NotZero(t, user.Username)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.PasswordChangedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	newUser := createRandomUser(t)
	user, err := testQueries.GetUser(context.Background(), newUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, newUser.Username, user.Username)
	require.Equal(t, newUser.HashedPassword, user.HashedPassword)
	require.Equal(t, newUser.Email, user.Email)
	require.Equal(t, newUser.FullName, user.FullName)
	require.Equal(t, newUser.CreatedAt, user.CreatedAt)
	require.Equal(t, newUser.PasswordChangedAt, user.PasswordChangedAt)
}
