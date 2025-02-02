package utils

import (
	"crypto/rand"
	"math/big"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateId() string {
	id, _ := generateRandomString(3)
	return id
}

func generateRandomString(length int) (string, error) {
	var result string
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result += string(charset[idx.Int64()])
	}
	return result, nil
}
