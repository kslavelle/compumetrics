package identity

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func hashPassword(password string) string {
	passwordHashingKey := os.Getenv("PASSWORD_HASH")
	hasher := hmac.New(sha256.New, []byte(passwordHashingKey))
	hasher.Write([]byte(password))

	return hex.EncodeToString(hasher.Sum(nil))
}
