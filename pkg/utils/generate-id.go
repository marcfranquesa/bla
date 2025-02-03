package utils

import (
	"crypto/sha256"
	"math/big"
	"strings"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateId(s string, length int) string {
	data := []byte(s)

	hash := sha256.New()
	hash.Write(data)
	hashedData := hash.Sum(nil)
	return encode(hashedData)[:length]
}

func encode(data []byte) string {
	var result strings.Builder
	num := new(big.Int).SetBytes(data)

	base := big.NewInt(62)
	zero := big.NewInt(0)

	for num.Cmp(zero) > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod)
		result.WriteByte(charset[mod.Int64()])
	}

	if result.Len() == 0 {
		result.WriteByte(charset[0])
	}

	return result.String()
}
