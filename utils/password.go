package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// Simple password hashing utility using SHA-256. For learn purpose
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// Compare a hashed password with a plain text password
func VerifyPassword(hashedPassword, password string) bool {
	return HashPassword(password) == hashedPassword
}
