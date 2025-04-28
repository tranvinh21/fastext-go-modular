package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) (string, error) {
	// Hash password using SHA256
	hashedPassword := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hashedPassword[:]), nil
}

func VerifyPassword(password string, hashedPassword string) bool {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return false
	}
	return hashedPassword == hashedPassword
}
