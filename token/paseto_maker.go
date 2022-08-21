package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto             *paseto.V2
	symmetricSecretKey []byte
}

func NewPasetoMaker(symmetricSecretKey string) (Maker, error) {
	if len(symmetricSecretKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("Invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	return &PasetoMaker{paseto: paseto.NewV2(), symmetricSecretKey: []byte(symmetricSecretKey)}, nil
}

func (pasetoMaker *PasetoMaker) GenerateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return pasetoMaker.paseto.Encrypt(pasetoMaker.symmetricSecretKey, payload, nil)
}

func (pasetoMaker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := pasetoMaker.paseto.Decrypt(token, pasetoMaker.symmetricSecretKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
