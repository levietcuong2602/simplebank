package test

import (
	"testing"

	"github.com/levietcuong2602/simplebank/utils"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := utils.RandomString(6)
	hashedPassword, err := utils.EncryptPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = utils.ComparePassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := utils.RandomString(8)
	err = utils.ComparePassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
