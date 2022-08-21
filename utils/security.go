package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword return hashed password
func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Fail hash password: %w", err)
	}

	return string(hashedPassword), err
}

// ComparePassword checks if provided password corrects or not
func ComparePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
