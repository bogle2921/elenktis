package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/argon2"
)

func GenerateSalt() []byte {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		log.Println(err)
		panic(err)
	}
	return salt
}

func HashPassword(salt []byte, pass string) string {
	hash := argon2.Key([]byte(pass), salt, 3, 32*1024, 4, 32)
	return base64.RawStdEncoding.EncodeToString(hash)
}

func VerifyHash(password string, hash string, salt []byte) bool {
	givenHash := argon2.Key([]byte(password), salt, 3, 32*1024, 4, 32)
	return base64.RawStdEncoding.EncodeToString(givenHash) == hash
}
