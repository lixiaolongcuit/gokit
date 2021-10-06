package cryptox

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Bytes(bytes []byte) string {
	sha256Bytes := sha256.Sum256(bytes)
	sha256Str := hex.EncodeToString(sha256Bytes[:])
	return sha256Str
}

func Sha256Str(str string) string {
	return Sha256Bytes([]byte(str))
}
