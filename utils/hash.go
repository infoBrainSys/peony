package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// NewHash 加密密码
func NewHash(pass []byte) (hPassword []byte, err error) {
	return bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
}

// NewDeHash 解密密码
func NewDeHash(hPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hPassword, password)
}

func NewHashPeony(peony string) string {
	hash := sha256.New()
	hash.Write([]byte(peony))
	hashStr := hash.Sum(nil)
	hashValue := hex.EncodeToString(hashStr)
	return hashValue
}
