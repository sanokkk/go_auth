package internal

import (
	"crypto/sha256"
	"encoding/hex"
)

type hasher interface {
	hash(password string) string
}

type Sha256Hash struct{}

func (hasher Sha256Hash) hash(password string) string {
	sh := sha256.Sum256([]byte(password))
	hashStr := hex.EncodeToString(sh[:])

	return hashStr
}

func HashPassword(hasher hasher, password string) string {
	return hasher.hash(password)
}
