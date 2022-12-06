package helpers

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())
	// Generate a random sequence of bytes
	b := make([]byte, 32)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	rnd.Read(b)

	// Encode the bytes as a hexadecimal string
	return hex.EncodeToString(b)
}
