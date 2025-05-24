package encrypt

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)

	return hex.EncodeToString(hashBytes)
}

func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}

func HashPassword(password, salt string) string {
	saltedPassword := password + salt

	hashPass := sha256.Sum256([]byte(saltedPassword))
	return hex.EncodeToString(hashPass[:])
}
