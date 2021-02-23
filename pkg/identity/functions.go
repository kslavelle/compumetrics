package identity

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"unicode"
)

func hashPassword(password string) string {
	passwordHashingKey := os.Getenv("PASSWORD_HASH")
	hasher := hmac.New(sha256.New, []byte(passwordHashingKey))
	hasher.Write([]byte(password))

	return hex.EncodeToString(hasher.Sum(nil))
}

func verifyPassword(s string) (eightOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			return false, false, false, false
		}
	}
	eightOrMore = letters >= 8
	return
}
