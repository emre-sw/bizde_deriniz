package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 6 digits random number
func GenerateVerificationToken() string {
	n, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%06d", n)
}
