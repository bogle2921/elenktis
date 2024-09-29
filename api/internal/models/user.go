package models

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/argon2"
)

type User struct {
	ID       int
	Admin    bool
	Username string `binding:"required"`
	Hash     string `binding:"required"`
	Salt     [16]byte
}

func generateSalt() []byte {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		log.Println(err)
		panic(err)
	}
	return salt //store in pg
}
func HashPassword(salt []byte, pass string) string {
	hash := argon2.Key([]byte(pass), salt, 3, 32*1024, 4, 32)
	return base64.RawStdEncoding.EncodeToString(hash)
}
