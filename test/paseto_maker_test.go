package test

import (
	"testing"
	"time"

	"github.com/levietcuong2602/simplebank/token"
	"github.com/levietcuong2602/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func TestPasetoMakerToken(t *testing.T) {
	maker, err := token.NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	username := utils.RandomString(16)
	duration := time.Minute
	issueAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, err := maker.GenerateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issueAt, payload.IssueAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredTokenPaseto(t *testing.T) {
	maker, err := token.NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	tokenPaseto, err := maker.GenerateToken(utils.RandomString(6), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, tokenPaseto)

	payload, err1 := maker.VerifyToken(tokenPaseto)
	require.Error(t, err1)
	require.EqualError(t, err1, token.ErrExpiredToken.Error())
	require.Nil(t, payload)
}
