package test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/levietcuong2602/simplebank/token"
	"github.com/levietcuong2602/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func TestJwtToken(t *testing.T) {
	maker, err := token.NewJwtMaker(utils.RandomString(32))
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

func TestInvalidToken(t *testing.T) {
	payload, err := token.NewPayload(utils.RandomString(32), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	tokenJwt, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := token.NewJwtMaker(utils.RandomString(32))
	require.NoError(t, err)

	payload, err1 := maker.VerifyToken(tokenJwt)
	require.Error(t, err1)
	require.EqualError(t, err1, token.ErrInvalidToken.Error())
	require.Nil(t, payload)

}
