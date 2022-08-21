package token

import "time"

type Maker interface {
	// GenerateToken generate a new token for specific username and duration
	GenerateToken(username string, duration time.Duration) (string, error)

	// VerifyToken check token is valid or not
	VerifyToken(token string) (*Payload, error)
}
